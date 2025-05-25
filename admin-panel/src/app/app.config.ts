import { APP_INITIALIZER, ApplicationConfig, provideZoneChangeDetection } from '@angular/core';
import { provideRouter } from '@angular/router';

import { routes } from './app.routes';
import { AuthInterceptor } from './interceptors/auth.interceptor';
import { HTTP_INTERCEPTORS, withInterceptorsFromDi, provideHttpClient } from '@angular/common/http';
import { of, tap } from 'rxjs';
import { catchError } from 'rxjs';
import { AuthService } from './services/auth.service';
import { Observable } from 'rxjs';

export function initializeAppFactory(authService: AuthService): () => Observable<any> {
  console.log('APP_INITIALIZER: Factory function called.'); // <--- ЛОГ 22
  return () => {
    console.log('APP_INITIALIZER: Inner function (returned by factory) EXECUTING - calling authService.checkSession()'); // <--- ЛОГ 23
    return authService.checkSession().pipe(
      tap(result => {
        console.log('APP_INITIALIZER: authService.checkSession() COMPLETED. Result:', result); // <--- ЛОГ 24 (покажет, что вернул checkSession, включая null из catchError)
      }),
      catchError(error => {
        // Этот catchError сработает, если checkSession() сам по себе вернет throwError,
        // но сейчас он возвращает of(null) в своем catchError, так что этот блок может не вызываться.
        console.error('APP_INITIALIZER: CRITICAL - Error during authService.checkSession() in initializeAppFactory:', error); // <--- ЛОГ 25
        return of(null); // Продолжаем запуск приложения
      })
    );
  }
}

export const appConfig: ApplicationConfig = {
  providers: [provideZoneChangeDetection({ eventCoalescing: true }), provideRouter(routes), provideHttpClient(withInterceptorsFromDi()),
    {
      provide: HTTP_INTERCEPTORS,
      useClass: AuthInterceptor,
      multi: true
    },
    {
      provide: APP_INITIALIZER,
      useFactory: initializeAppFactory,
      deps: [AuthService],
      multi: true
    }
  ]
};
