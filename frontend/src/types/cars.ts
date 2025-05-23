export interface Car {
    id: number,
    brand_id: number,
    model: string,
    name: string,
    year: number,
    price_per_hour: number,
    image_url: string,
    available: boolean,
    created_at: string,
    updated_at: string
}

export interface CarsResponse {
    cars: Car[]
}