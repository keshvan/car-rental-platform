import { Routes } from '@angular/router';
import { authGuard } from './gurads/auth.guard';
import { LoginComponent } from './pages/login/login.component';
import { LayoutComponent } from './components/layout/layout.component';
import { CarManagementComponent } from './pages/car-management/car-management.component';
import { RentManagementComponent } from './pages/rent-management/rent-management.component';
import { UserManagementComponent } from './pages/user-management/user-management.component';
import { ReviewManagementComponent } from './pages/review-management/review-management.component';

export const routes: Routes = [
    {
        path: 'login',
        component: LoginComponent
    },
    {
        path: '',
        component: LayoutComponent,
        canActivate: [authGuard],
        children: [
            { path: '', redirectTo: 'cars', pathMatch: 'full' },
            { path: 'cars', component: CarManagementComponent },
            { path: 'rents', component: RentManagementComponent },
            { path: 'users', component: UserManagementComponent },
            { path: 'reviews', component: ReviewManagementComponent }
        ]
    },
    { path: '**', redirectTo: 'login' }
];
