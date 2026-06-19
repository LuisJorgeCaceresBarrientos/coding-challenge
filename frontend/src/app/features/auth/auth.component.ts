import { Component, Output, EventEmitter, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { ApiService } from '../../core/services/api.service';

@Component({
  selector: 'app-auth',
  standalone: true,
  imports: [CommonModule, FormsModule],
  templateUrl: './auth.component.html'
})
export class AuthComponent {
  private apiService = inject(ApiService);

  // Evento que emitimos cuando el usuario se autentica exitosamente
  @Output() authSuccess = new EventEmitter<void>();

  // Estado local del componente
  currentView: 'login' | 'register' = 'login';
  authUsername = '';
  authPassword = '';
  errorMessage = '';
  successMessage = '';

  /**
   * Cambia la vista entre login y registro, limpiando errores previos.
   */
  setView(view: 'login' | 'register'): void {
    this.currentView = view;
    this.errorMessage = '';
    this.successMessage = '';
    this.authUsername = '';
    this.authPassword = '';
  }

  /**
   * Ejecuta el flujo de inicio de sesión.
   */
  onLogin(): void {
    this.errorMessage = '';
    this.successMessage = '';

    if (!this.authUsername || !this.authPassword) {
      this.errorMessage = 'Por favor, completa todos los campos.';
      return;
    }

    this.apiService.login(this.authUsername, this.authPassword).subscribe({
      next: (res) => {
        this.apiService.saveToken(res.token);
        this.authSuccess.emit(); // Notificamos al componente padre
      },
      error: (err) => {
        this.errorMessage = err.error?.error || 'Usuario o contraseña incorrectos.';
      }
    });
  }

  /**
   * Ejecuta el flujo de registro.
   */
  onRegister(): void {
    this.errorMessage = '';
    this.successMessage = '';

    if (!this.authUsername || !this.authPassword) {
      this.errorMessage = 'Por favor, completa todos los campos.';
      return;
    }

    this.apiService.register(this.authUsername, this.authPassword).subscribe({
      next: () => {
        this.successMessage = 'Registro exitoso. Ahora puedes iniciar sesión.';
        this.setView('login');
      },
      error: (err) => {
        this.errorMessage = err.error?.error || 'Ocurrió un error al registrar el usuario.';
      }
    });
  }
}
