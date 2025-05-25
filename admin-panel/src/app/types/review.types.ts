export interface Review {
    id: number;         
    rent_id: number;    
    user_id: number;    
    email: string;
    car_id: number;     
    rating: number;     
    comment?: string;    
    created_at: string; 
}

export interface ReviewsResponse {
    reviews: Review[];
}