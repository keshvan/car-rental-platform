import { Component, OnInit } from '@angular/core';
import { CommonModule, DatePipe } from '@angular/common';
import { Rent } from '../../types/rent.types'; // Убедитесь, что Rent импортирован
import { RentService } from '../../services/rent.service';
import { HttpErrorResponse } from '@angular/common/http';

@Component({
  selector: 'app-rent-management',
  standalone: true,
  imports: [CommonModule, DatePipe],
  templateUrl: './rent-management.component.html',
})
export class RentManagementComponent implements OnInit {
  rents: Rent[] = [];
  isLoading = false;
  loadErrorMessage: string | null = null;
  actionSuccessMessage: string | null = null;
  actionErrorMessage: string | null = null;

  constructor(private rentService: RentService) {}

  ngOnInit(): void {
    this.loadRents();
  }

  loadRents(): void {
    this.isLoading = true;
    this.loadErrorMessage = null;
    this.clearActionMessages();

    this.rentService.getAllRents().subscribe({
      next: (response) => {
        this.rents = response.rents;
        this.isLoading = false;
      },
      error: (error: HttpErrorResponse) => {
        this.loadErrorMessage = 'Ошибка загрузки аренд';
        this.isLoading = false;
        console.error('Error loading rents:', error);
      }
    });
  }

  private clearActionMessages(): void {
    this.actionSuccessMessage = null;
    this.actionErrorMessage = null;
  }

  confirmCancelRent(rent: Rent): void {
    if (rent.status && (rent.status.toLowerCase() === 'cancelled' || rent.status.toLowerCase() === 'completed')) {
        this.actionErrorMessage = `Аренду ID ${rent.id} нельзя отменить, так как она уже '${rent.status}'.`;
        setTimeout(() => this.clearActionMessages(), 4000);
        return;
    }

    this.clearActionMessages();
    const confirmation = window.confirm(
      `Вы уверены, что хотите ОТМЕНИТЬ аренду ID: ${rent.id} для машины "${rent.car_name}"?`
    );

    if (confirmation) {
      this.rentService.cancelRent(rent.id).subscribe({
        next: () => {
          this.actionSuccessMessage = `Аренда ID: ${rent.id} успешно отменена.`;
          this.loadRents();
          setTimeout(() => this.clearActionMessages(), 3000);
        },
        error: (error: HttpErrorResponse) => {
          this.actionErrorMessage = `Ошибка отмены аренды ID: ${rent.id}. Статус: ${error.status}. ${error.error?.message || error.message}`;
          console.error('Error cancelling rent:', error);
          setTimeout(() => this.clearActionMessages(), 5000);
        }
      });
    }
  }

  confirmDeleteRent(rent: Rent): void {
    this.clearActionMessages();
    const confirmation = window.confirm(
      `Вы уверены, что хотите УДАЛИТЬ аренду ID: ${rent.id} (машина: "${rent.car_name}")? Это действие необратимо.`
    );

    if (confirmation) {
      this.rentService.deleteRent(rent.id).subscribe({
        next: () => {
          this.actionSuccessMessage = `Аренда ID: ${rent.id} успешно удалена.`;
          this.loadRents()
          setTimeout(() => this.clearActionMessages(), 3000);
        },
        error: (error: HttpErrorResponse) => {
          this.actionErrorMessage = `Ошибка удаления аренды ID: ${rent.id}. Статус: ${error.status}. ${error.error?.message || error.message}`;
          console.error('Error deleting rent:', error);
          setTimeout(() => this.clearActionMessages(), 5000);
        }
      });
    }
  }

  public getStatusClass(status: string | undefined | null): object {
    const lowerStatus = status?.toLowerCase();

    return {
      'bg-green-100 text-green-800': lowerStatus === 'completed',
      'bg-blue-100 text-blue-800': lowerStatus === 'active',
      'bg-orange-100 text-orange-800': lowerStatus === 'cancelled',
    };
  }
}