import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { User, UsersResponse, UpdateUserRolePayload } from '../types/user.types'; // Предполагаем, что типы здесь
import { environment } from '../../environments/environment'; // Для URL API

@Injectable({
  providedIn: 'root'
})
export class UserService {
  private apiUrl = `${environment.authUrl}users`;

  constructor(private http: HttpClient) { }

  getAllUsers(): Observable<UsersResponse> {
    return this.http.get<UsersResponse>(this.apiUrl, { withCredentials: true });
  }

  updateUserRole(userId: number, payload: UpdateUserRolePayload): Observable<User> {
    return this.http.patch<User>(`${this.apiUrl}/${userId}`, payload, { withCredentials: true });
  }

  deleteUser(userId: number): Observable<void> {
    return this.http.delete<void>(`${this.apiUrl}/${userId}`, { withCredentials: true });
  }
}