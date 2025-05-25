import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from '../../environments/environment';
import { ReviewsResponse } from '../types/review.types';

@Injectable({
  providedIn: 'root'
})
export class ReviewService {
  private readonly apiUrl = environment.carUrl;

  constructor(private http: HttpClient) { }

  getReviews(): Observable<ReviewsResponse> {
    return this.http.get<ReviewsResponse>(`${this.apiUrl}reviews`);
  }

  deleteReview(reviewId: number): Observable<void> {
    return this.http.delete<void>(`${this.apiUrl}reviews/${reviewId}`);
  }

  verifyReview(reviewId: number): Observable<void> {
    return this.http.patch<void>(`${this.apiUrl}reviews/${reviewId}`, {});
  }
}