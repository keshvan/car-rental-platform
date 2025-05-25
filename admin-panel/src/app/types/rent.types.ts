export interface Rent {
    id: number;           
    user_id: number;      
    car_id: number;       
    car_name: string;
    price_per_hour: number; 
    start_date: string;   
    end_date: string;     
    total_price?: number | null; 
    status: string;       
    created_at: string;   
    updated_at: string;   
}

export interface RentsResponse {
    rents: Rent[];
}