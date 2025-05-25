export interface Car {
    id: number;
    brand_id: number; 
    model: string;
    name: string; 
    year: number;
    price_per_hour: number;
    image_url: string;
    available: boolean;
    created_at: string;
    updated_at: string;
}

export interface Brand {
    id: number;
    name: string;
}

export interface CarsResponse {
    cars: Car[];
}

export interface BrandsResponse {
    brands: Brand[];
}

export interface NewCarRequest {
    brand_id: number;
    model: string;
    year: number;
    price_per_hour: number;
    image_url: string | undefined;
}

export interface UpdateCarRequest {
    brand_id: number;
    model: string;
    year: number;
    price_per_hour: number;
    image_url: string | undefined;
    available: boolean;
}
