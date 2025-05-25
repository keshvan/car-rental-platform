import { Component, OnInit, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HttpClientModule } from '@angular/common/http'; // Если сервис не в standalone
import { UserService } from '../../services/user.service';
import { User } from '../../types/user.types';
import { AuthService } from '../../services/auth.service'; // Для получения текущего пользователя

@Component({
  selector: 'app-user-management',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './user-management.component.html',
})
export class UserManagementComponent implements OnInit {

  constructor(private userService: UserService, private authService: AuthService) {}

  users: User[] = [];
  isLoading = true;
  error: string | null = null;
  currentUserId: number | null = null;

  ngOnInit(): void {
    this.authService.currentUser.subscribe(user => {
      this.currentUserId = user ? user.id : null;
    });
    this.loadUsers();
  }

  loadUsers(): void {
    this.isLoading = true;
    this.error = null;
    this.userService.getAllUsers().subscribe({
      next: (response) => {
        this.users = response.users;
        this.isLoading = false;
      },
      error: (err) => {
        this.error = 'Не удалось загрузить список пользователей. ' + (err.error?.message || err.message);
        this.isLoading = false;
      }
    });
  }

  promoteToAdmin(userId: number): void {
    if (userId === this.currentUserId) {
      alert('Вы не можете изменить роль самому себе.');
      return;
    }
    if (confirm('Вы уверены, что хотите повысить этого пользователя до администратора?')) {
      this.userService.updateUserRole(userId, { role: 'admin' }).subscribe({
        next: (updatedUser) => {
          this.loadUsers();
          alert('Роль пользователя успешно обновлена.');
        },
        error: (err) => {
          alert('Ошибка при обновлении роли пользователя. ' + (err.error?.message || err.message));
        }
      });
    }
  }

  deleteUser(userId: number): void {
    if (userId === this.currentUserId) {
      alert('Вы не можете удалить самого себя.');
      return;
    }
    if (confirm('Вы уверены, что хотите удалить этого пользователя? Это действие необратимо.')) {
      this.userService.deleteUser(userId).subscribe({
        next: () => {
          this.loadUsers();
          alert('Пользователь успешно удален.');
        },
        error: (err) => {
          alert('Ошибка при удалении пользователя. ' + (err.error?.message || err.message));
        }
      });
    }
  }
}