import { Injectable } from "@angular/core";
import { environment } from "../../environments/environment";
import { HttpClient } from "@angular/common/http";
import { Router } from "@angular/router";
import { BehaviorSubject, catchError, filter, Observable, of, switchMap, take, tap, throwError } from "rxjs";


export interface User {
    id: number;
    email: string;
    role: string;
}

interface SessionCheckResponse {
    user: User | null;
    is_active: boolean;
}

interface LoginResponse {
    access_token: string;
    user: User;
}

interface RefreshResponse {
    access_token: string;
}


@Injectable({
    providedIn: 'root'
})
export class AuthService {
    private readonly authUrl = environment.authUrl;
    private readonly carUrl = environment.carUrl;

    private currentUserSource = new BehaviorSubject<User | null>(null);
    currentUser = this.currentUserSource.asObservable();

    private isAuthenticatedSource = new BehaviorSubject<boolean>(false);
    isAuthenticated = this.isAuthenticatedSource.asObservable();

    private isRefreshingToken = false;
    private tokenRefreshed = new BehaviorSubject<boolean | null>(null);

    constructor(private http: HttpClient, private router: Router) {}

    checkSession() {
        console.log('AuthService: checkSession() called'); // <--- ЛОГ 2
        return this.http.get<SessionCheckResponse>(`${this.authUrl}check-session`, { withCredentials: true }).pipe(
            tap((response) => {
                console.log('AuthService: checkSession() - response received:', response);
                if (response.is_active && response.user) {
                    this.currentUserSource.next(response.user);
                    this.isAuthenticatedSource.next(true);
                    console.log('AuthService: checkSession() - session active. isAuthenticatedSource set to true. User:', response.user);
                } else {
                    console.log('AuthService: checkSession() - session NOT active or no user.');
                    this.clearSession();
                }
            }),
            catchError((error) => {
                console.error('AuthService: checkSession() - FAILED:', error); // <--- ЛОГ 6
                this.clearSession();
                return throwError(() => new Error('Session check failed'));
            })
        )
    }

    login(email: string, password: string): Observable<LoginResponse> {
        console.log('AuthService: login() called for email:', email); // 
        return this.http.post<LoginResponse>(`${this.authUrl}login`, { email, password }, { withCredentials: true }).pipe(
            tap((response) => {
                console.log('AuthService: login() - response received:', response);
                if (response.user && response.access_token) {
                    localStorage.setItem('access_token', response.access_token);
                    this.currentUserSource.next(response.user);
                    this.isAuthenticatedSource.next(true);
                    console.log('AuthService: login() - successful. isAuthenticatedSource set to true.');
                }
            }),
            catchError((error) => {
                console.error('AuthService: login() - FAILED:', error);
                return throwError(() => new Error('Login failed'));
            })
        )
    }

    logout() {
        console.log('AuthService: logout() called');
        return this.http.post(`${this.authUrl}logout`, {}, { withCredentials: true }).pipe(
            tap(() => {
                console.log('AuthService: logout() - successful response.');
                this.clearSession();
            }),
            catchError(() => {
                console.warn('AuthService: logout() - FAILED. Clearing session locally.'); // <--- ЛОГ 13
                this.clearSession();
                return throwError(() => new Error('Logout failed'));
            })
        )
    }

    refresh(): Observable<RefreshResponse> {
        if (this.isRefreshingToken) {
            return this.tokenRefreshed.pipe(
                filter((isRefreshed) => isRefreshed !== null),
                take(1),
                switchMap((res) => {
                    if (res) {
                        const token = localStorage.getItem('access_token');
                        return token ? of({access_token: token}) : throwError(() => new Error('No token found'));
                    }
                    return throwError(() => new Error('Token refresh failed'));
                })
            ) as Observable<RefreshResponse>;
        }

        this.isRefreshingToken = true;
        this.tokenRefreshed.next(null);

        return this.http.post<RefreshResponse>(`${this.authUrl}refresh`, {}, { withCredentials: true }).pipe(
            tap((response) => {
                if (response.access_token) {
                    localStorage.setItem('access_token', response.access_token);
                    this.isRefreshingToken = false;
                    this.tokenRefreshed.next(true);
                }
            }),
            catchError(() => {
                this.isRefreshingToken = false;
                this.tokenRefreshed.next(false);
                this.logoutAndRedirect();
                return throwError(() => new Error('Token refresh failed'));
            })
        )
        
    }

    getAccessToken(): string | null {
        return localStorage.getItem("access_token");
    }

    isLoggedIn(): boolean {
        console.log('AuthService: isLoggedIn() called. Returning:', this.isAuthenticatedSource.value); // <--- ЛОГ 14 (если используется)
        return this.isAuthenticatedSource.value;
    }
    
    logoutAndRedirect(): void {
        console.log('AuthService: logoutAndRedirect() called.');
        this.clearSession();
        this.router.navigate(['/login']);
    }

    private clearSession() {
        localStorage.removeItem('access_token');
        this.currentUserSource.next(null);
        this.isAuthenticatedSource.next(false);
    }
}
