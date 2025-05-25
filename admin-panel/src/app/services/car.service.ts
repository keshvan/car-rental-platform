import { Injectable } from "@angular/core";
import { environment } from "../../environments/environment";
import { Observable } from "rxjs";
import { BrandsResponse, CarsResponse, NewCarRequest, UpdateCarRequest } from "../types/car.types";
import { HttpClient } from "@angular/common/http";


@Injectable({
    providedIn: 'root'
})

export class CarService {
    private carsUrl = environment.carUrl;

    constructor(private http: HttpClient) {}

    getCars(): Observable<CarsResponse> {
        return this.http.get<CarsResponse>(`${this.carsUrl}cars`);
    }

    getBrands(): Observable<BrandsResponse> {
        return this.http.get<BrandsResponse>(`${this.carsUrl}brands`);
    }

    deleteCar(carId: number): Observable<void> {
        return this.http.delete<void>(`${this.carsUrl}cars/${carId}`);
    }

    newCar(newCarRequest: NewCarRequest): Observable<void> {
        return this.http.post<void>(`${this.carsUrl}cars`, newCarRequest);
    }

    updateCar(carId: number, updateCarRequest: UpdateCarRequest): Observable<void> {
        return this.http.patch<void>(`${this.carsUrl}cars/${carId}`, updateCarRequest);
    }
}