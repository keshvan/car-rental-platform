import { carApi } from "../api/api";
import type { CarsResponse, Car } from "../types/cars";

export const carService = {
    getCars: async (): Promise<Car[]> => {
        try {
            const res = await carApi.get<CarsResponse>('/cars')
            return res.data.cars
        } catch (error) {
            throw error
        }
    }
}