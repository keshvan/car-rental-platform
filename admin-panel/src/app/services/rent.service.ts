// admin-panel/src/app/services/rent.service.ts
import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from '../../environments/environment';
import { RentsResponse, Rent } from '../types/rent.types'; // Добавлен Rent

@Injectable({
    providedIn: 'root'
})
export class RentService {
    private readonly apiUrl = environment.carUrl;

    constructor(private http: HttpClient) { }

    getAllRents(): Observable<RentsResponse> {
        return this.http.get<RentsResponse>(`${this.apiUrl}rents`);
    }

    cancelRent(rentId: number): Observable<void> {
        return this.http.patch<void>(`${this.apiUrl}rents/cancel/${rentId}`, {});
    }

    deleteRent(rentId: number): Observable<void> {
        return this.http.delete<void>(`${this.apiUrl}rents/${rentId}`);
    }
}