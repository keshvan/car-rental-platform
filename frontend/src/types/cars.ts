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

export interface Brand {
    id: number,
    name: string
}

export interface CarsResponse {
    cars: Car[]
}

export interface BrandsRespone {
    brands: Brand[]
}

export interface Balance {
    id: number,
    amount: number
}

export interface Rent {
    id: number,
    user_id: number,
    car_id: number,
    car_name: string,
    price_per_hour: number,
    start_date: string,
    end_date: string,
    total_price: number,
    status: string,
    created_at: string,
    updated_at: string
}

export interface RentsResponse {
    rents: Rent[]
}

export interface Review {
    id: number,
    car_id: number,
    email: string,
    rating: number,
    comment: string,
    created_at: string
}

export interface CarWithReviews {
    car: Car,
    reviews: Review[]
}

