import { inject } from "@angular/core";
import { AuthService } from "../services/auth.service";
import { CanActivateFn, ActivatedRouteSnapshot, Router, RouterStateSnapshot } from "@angular/router";
import { map, take, tap } from "rxjs";


export const authGuard: CanActivateFn = (route: ActivatedRouteSnapshot, state: RouterStateSnapshot) => {

    const authService = inject(AuthService);
    const router = inject(Router);

    console.log('AuthGuard: Executing for route:', state.url); // <--- ЛОГ 18


    return authService.isAuthenticated.pipe(
        take(1),
        map((isAuthenticated) => {
            console.log(`AuthGuard: isAuthenticated value from authService.isAuthenticated (after take(1)): ${isAuthenticated}`); // <--- ЛОГ 19
            if (isAuthenticated) {
                console.log('AuthGuard: Access GRANTED for route:', state.url); // <--- ЛОГ 20
                return true;
            }
            console.log('AuthGuard: Access DENIED for route:', state.url, '- Redirecting to /login'); // <--- ЛОГ 21
            router.navigate(['/login']);
            return false;
        })
    );
};
