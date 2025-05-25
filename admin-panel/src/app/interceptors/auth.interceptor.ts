import { HttpErrorResponse, HttpEvent } from "@angular/common/http";
import { HttpHandler } from "@angular/common/http";
import { catchError, filter, switchMap, take, throwError } from "rxjs";
import { HttpRequest } from "@angular/common/http";
import { inject, Injectable } from "@angular/core";
import { AuthService } from "../services/auth.service";
import { HttpInterceptor } from "@angular/common/http";
import { BehaviorSubject, Observable } from "rxjs";
import { Router } from "@angular/router";

@Injectable()
export class AuthInterceptor implements HttpInterceptor {
  private isRefreshing = false;
  private refreshTokenSubject: BehaviorSubject<any> = new BehaviorSubject<any>(null);

  constructor(private authService: AuthService, private router: Router) {}

  intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    const accessToken = this.authService.getAccessToken();
    
    if (accessToken) {
      req = this.addTokenToRequest(req, accessToken);
    }

    return next.handle(req).pipe(
      catchError(error => {
        if (error instanceof HttpErrorResponse && error.status === 401) {
          if (req.url.includes('/refresh')) {
            this.authService.logoutAndRedirect();
            return throwError(() => error);
          }
          return this.handle401Error(req, next);
        }
        return throwError(() => error);
      })
    );
  }

  private addTokenToRequest(request: HttpRequest<any>, token: string): HttpRequest<any> {
    return request.clone({
      setHeaders: {
        Authorization: `Bearer ${token}`
      }
    });
  }

  private handle401Error(request: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    if (!this.isRefreshing) {
      this.isRefreshing = true;
      this.refreshTokenSubject.next(null);

      return this.authService.refresh().pipe(
        switchMap((tokenResponse) => {
          this.isRefreshing = false;
          this.refreshTokenSubject.next(tokenResponse.access_token);
          return next.handle(this.addTokenToRequest(request, tokenResponse.access_token));
        }),
        catchError((err) => {
          this.isRefreshing = false;
          this.authService.logoutAndRedirect();
          return throwError(() => err);
        })
      );
    } else {
      return this.refreshTokenSubject.pipe(
        filter(token => token != null),
        take(1),
        switchMap(jwt => next.handle(this.addTokenToRequest(request, jwt)))
      );
    }
  }
}