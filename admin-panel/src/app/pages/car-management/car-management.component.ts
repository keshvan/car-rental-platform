import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { CarService } from '../../services/car.service';
import { Brand, Car, NewCarRequest, UpdateCarRequest } from '../../types/car.types';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { FormsModule, NgForm } from '@angular/forms';
import { HttpErrorResponse } from '@angular/common/http';

interface CarFormModel {
  brand_id: number | null;
  model: string;
  year: number | null;
  price_per_hour: number | null;
  image_url: string | undefined;
  available: boolean;
}


@Component({
  selector: 'app-car-management',
  standalone: true,
  imports: [CommonModule, MatProgressSpinnerModule, FormsModule],
  templateUrl: './car-management.component.html',
  styleUrl: './car-management.component.css'
})
export class CarManagementComponent implements OnInit {
  allCars: Car[] = [];
  displayedCars: Car[] = [];
  brands: Brand[] = [];
  isLoadingCars = false;
  isLoadingBrands = false;

  carsErrorMessage: string | null = null;
  brandsErrorMessage: string | null = null;

  selectedBrandId: number | string = "";

  newCarErrorMessage: string | null = null;
  newCarSuccessMessage: string | null = null;
  updateCarErrorMessage: string | null = null;
  updateCarSuccessMessage: string | null = null;


  showAddCarForm = false;
  newCarForm: CarFormModel = this.getInitialCarFormModel();
  isSubmittingCar = false;

  showEditCarForm = false;
  editingCarFormModel: CarFormModel = this.getInitialCarFormModel();
  editingCarId: number | null = null;
  isSubmittingEditCar = false;

  constructor(private carService: CarService) {}

  ngOnInit() {
    this.loadCars();
    this.loadBrands();
  }

  private getInitialCarFormModel(): CarFormModel {
    return {
      brand_id: null,
      model: '',
      year: null,
      price_per_hour: null,
      image_url: '',
      available: true
    }
  }

  loadCars() {
    this.isLoadingCars = true;
    this.carService.getCars().subscribe({
      next: (cars) => {
        this.allCars = cars.cars;
        this.applyFilters();
        this.isLoadingCars = false;
      },
      error: (error) => {
        this.carsErrorMessage = 'Ошибка при загрузке автомобилей';
        this.isLoadingCars = false;
      }
    })
  }

  loadBrands() {
    this.isLoadingBrands = true;
    this.carService.getBrands().subscribe({
      next: (brands) => {
        this.brands = brands.brands;
        this.isLoadingBrands = false;
      },
      error: (error) => {
        this.brandsErrorMessage = 'Ошибка при загрузке брендов';
        this.isLoadingBrands = false;
      }
    })
  }

  applyFilters(): void {
    if (!this.selectedBrandId) {
      this.displayedCars = [...this.allCars];
    } else {
      this.displayedCars = this.allCars.filter(
        (car: Car) => car.brand_id === Number(this.selectedBrandId)
      );
    }
  }

  onBrandSelectionChange(): void {
    this.applyFilters();
  }

  onAddCarSubmit(form: NgForm): void {
    if (form.invalid || this.newCarForm.brand_id === null || this.newCarForm.year === null || this.newCarForm.price_per_hour === null) {
      this.newCarErrorMessage = "Пожалуйста, заполните все обязательные";
      Object.values(form.controls).forEach(control => {
        control.markAsTouched();
      });
      return;
    }
    this.newCarErrorMessage = null;
    this.newCarSuccessMessage = null;
    this.isSubmittingCar = true;

    const carDataToSend: NewCarRequest = {
      brand_id: Number(this.newCarForm.brand_id),
      model: this.newCarForm.model,
      year: this.newCarForm.year,
      price_per_hour: this.newCarForm.price_per_hour,
      image_url: this.newCarForm.image_url || undefined,
    };

    console.log(carDataToSend);

    this.carService.newCar(carDataToSend).subscribe({
      next: () => {
        this.newCarSuccessMessage = 'Автомобиль успешно добавлен.';
        this.isSubmittingCar = false;
        this.showAddCarForm = false;
        this.newCarForm = this.getInitialCarFormModel();
        form.resetForm(this.getInitialCarFormModel());
        this.loadCars();

        setTimeout(() => this.newCarSuccessMessage = null, 3000);
      },
      error: (error: HttpErrorResponse) => {
        this.newCarErrorMessage = `Ошибка добавления автомобиля. Статус: ${error.status}. ${error.error?.message || error.message || 'Неизвестная ошибка бэкенда'}`;
        this.isSubmittingCar = false;
        console.error("Error adding car:", error);
      }
    });
  }

  toggleAddCarForm(): void {
    this.showAddCarForm = !this.showAddCarForm;
    if (this.showAddCarForm) {
      this.closeEditCarForm();
      this.newCarForm = this.getInitialCarFormModel();
      this.newCarErrorMessage = null;
      this.newCarSuccessMessage = null;
    }
  }

  closeAddCarForm(): void {
    this.showAddCarForm = false;
  }

  openEditCarForm(carToEdit: Car): void {
    this.closeAddCarForm();
    this.editingCarId = carToEdit.id;
    this.editingCarFormModel = {
        ...carToEdit, 
        image_url: carToEdit.image_url || ''
    };
    this.showEditCarForm = true;
    this.updateCarErrorMessage = null;
    this.updateCarSuccessMessage = null;
  }

  closeEditCarForm(): void {
    this.showEditCarForm = false;
    this.editingCarId = null;
  }

  onEditCarSubmit(form: NgForm): void {
    if (!this.editingCarId || form.invalid || this.editingCarFormModel.brand_id === null || this.editingCarFormModel.year === null || this.editingCarFormModel.price_per_hour === null) {
      this.updateCarErrorMessage = "Пожалуйста, заполните все обязательные поля корректно.";
       Object.values(form.controls).forEach(control => control.markAsTouched());
      return;
    }
    this.updateCarErrorMessage = null;
    this.updateCarSuccessMessage = null;
    this.isSubmittingEditCar = true;

    const carDataToUpdate: UpdateCarRequest = {
      brand_id: this.editingCarFormModel.brand_id,
      model: this.editingCarFormModel.model,
      year: this.editingCarFormModel.year,
      price_per_hour: this.editingCarFormModel.price_per_hour,
      image_url: this.editingCarFormModel.image_url || undefined,
      available: this.editingCarFormModel.available
    };

    this.carService.updateCar(this.editingCarId, carDataToUpdate).subscribe({
      next: () => {
        this.loadCars();
        this.isSubmittingEditCar = false;
        this.closeEditCarForm();
        setTimeout(() => this.updateCarSuccessMessage = null, 3000);
      },
      error: (error: HttpErrorResponse) => {
        this.updateCarErrorMessage = `Ошибка обновления. ${error.error?.message || error.message}`;
        this.isSubmittingEditCar = false;
      }
    });
  }

  deleteCar(carToDelete: Car): void {
    const confirmation = window.confirm(`Вы уверены, что хотите удалить автомобиль "${carToDelete.name}"?`);
    if (confirmation) {
      this.carService.deleteCar(carToDelete.id).subscribe({
        next: () => {
          this.loadCars();
        },
        error: (error) => {
          this.carsErrorMessage = 'Ошибка при удалении автомобиля';
        }
      })
    }
  }

  currentYear(): number {
    return new Date().getFullYear();
  }
}