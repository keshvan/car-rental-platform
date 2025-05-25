import { Component, OnInit } from '@angular/core';
import { CommonModule, DatePipe } from '@angular/common'; // DatePipe для форматирования даты
import { Review } from '../../types/review.types';
import { ReviewService } from '../../services/review.service';
import { HttpErrorResponse } from '@angular/common/http';

@Component({
  selector: 'app-review-management',
  standalone: true,
  imports: [CommonModule, DatePipe],
  templateUrl: './review-management.component.html',
})
export class ReviewManagementComponent implements OnInit {
  reviews: Review[] = [];
  isLoading = false;
  errorMessage: string | null = null;
  deleteSuccessMessage: string | null = null;
  deleteErrorMessage: string | null = null;

  verifySuccessMessage: string | null = null;
  verifyErrorMessage: string | null = null;

  constructor(private reviewService: ReviewService) { }

  ngOnInit(): void {
    this.loadReviews();
  }

  loadReviews(): void {
    this.isLoading = true;
    this.errorMessage = null;
    this.deleteSuccessMessage = null;
    this.deleteErrorMessage = null;

    this.reviewService.getReviews().subscribe({
      next: (response) => {
        this.reviews = response.reviews;
        console.log(response.reviews);
        this.isLoading = false;
      },
      error: (error: HttpErrorResponse) => {
        this.errorMessage = `Ошибка загрузки отзывов`;
        this.isLoading = false;
      }
    });
  }

  confirmDeleteReview(review: Review): void {
    this.deleteSuccessMessage = null;
    this.deleteErrorMessage = null;

    const confirmation = window.confirm(
      `Вы уверены, что хотите удалить отзыв ID: ${review.id} от пользователя ${review.email} (Рейтинг: ${review.rating})?\nКомментарий: "${review.comment || 'Нет комментария'}"`
    );

    if (confirmation) {
      this.reviewService.deleteReview(review.id).subscribe({
        next: () => {
          this.deleteSuccessMessage = `Отзыв ID: ${review.id} успешно удален.`;
          this.loadReviews();
          setTimeout(() => this.deleteSuccessMessage = null, 5000);
        },
        error: (error: HttpErrorResponse) => {
          this.deleteErrorMessage = `Ошибка удаления отзыва ID: ${review.id}. Статус: ${error.status}. ${error.error?.message || error.message}`;
          setTimeout(() => this.deleteErrorMessage = null, 5000);
        }
      });
    }
  }

  confirmApproveReview(review: Review): void {
    const confirmation = window.confirm(
      `Вы уверены, что хотите ОДОБРИТЬ отзыв ID: ${review.id} от пользователя ${review.email}?`
    );

    if (confirmation) {
      this.reviewService.verifyReview(review.id).subscribe({
        next: () => {
          this.verifySuccessMessage = `Отзыв ID: ${review.id} успешно одобрен.`;
          this.loadReviews();
        },
        error: (error: HttpErrorResponse) => {
          this.verifyErrorMessage = `Ошибка одобрения отзыва ID: ${review.id}. Статус: ${error.status}. ${error.error?.message || error.message}`;
          setTimeout(() => this.verifyErrorMessage = null, 5000);
        }
      });
    }
  }
}