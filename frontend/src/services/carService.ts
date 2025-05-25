import { carApi } from "../api/api";
import type { CarsResponse, Car, Brand, BrandsRespone, Balance, Rent, RentsResponse, CarWithReviews } from "../types/cars";

export const carService = {
    getCars: async (): Promise<Car[]> => {
        try {
            const res = await carApi.get<CarsResponse>('/cars')
            return res.data.cars
        } catch (error) {
            throw error
        }
    },

    getBrands: async (): Promise<Brand[]> => {
        try {
            const res = await carApi.get<BrandsRespone>('/brands')
            return res.data.brands
        } catch (error) {
            throw error
        }
    },

    updateBalance: async (id: number, amount: number): Promise<Balance> => {
        try {
            const res = await carApi.patch<Balance>(`/users/${id}/balance`, { amount })
            return res.data
        } catch (error) {
            throw error
        }
    },

    newRent: async (carId: number, startDate: string): Promise<{success: boolean}> => {
        try {
            await carApi.post('/rents', { car_id: carId, start_date: startDate })
            return {success: true}
        } catch (error) {
            throw error
        }
    },

    getMyRents: async (): Promise<Rent[]> => {
        try {
            const res = await carApi.get<RentsResponse>('/rents/me')
            return res.data.rents
        } catch (error) {
            throw error
        }
    },

    getRents: async (): Promise<Rent[]> => {
        try {
            const res = await carApi.get<RentsResponse>('/rents')
            return res.data.rents
        } catch (error) {
            throw error
        }
    },

    completeRent: async(rentId: number, rating: number, comment: string): Promise<{success: boolean}> => {
        try {
            await carApi.patch(`/rents/complete/${rentId}`, {rating, comment})
            return {success: true}
        } catch (error) {
            throw error
        }
    },

    getCarWithReviews: async (id: string | undefined): Promise<CarWithReviews> => {
        if (!id) {
            throw new Error('Car ID is required')
        }
        try {
            const res = await carApi.get<CarWithReviews>(`/cars/${id}`)
            return res.data
        } catch (error) {
            throw error
        }
    }
}