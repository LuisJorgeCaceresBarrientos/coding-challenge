import { Component, Output, EventEmitter, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { ApiService } from '../../core/services/api.service';
import Swal from 'sweetalert2';

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

  /**
   * Cambia la vista entre login y registro, limpiando errores previos.
   */
  setView(view: 'login' | 'register'): void {
    this.currentView = view;
    this.authUsername = '';
    this.authPassword = '';
  }

  /**
   * Ejecuta el flujo de inicio de sesión.
   */
  onLogin(): void {
    if (!this.authUsername || !this.authPassword) {
      Swal.fire({ icon: 'warning', title: 'Campos vacíos', text: 'Por favor, completa todos los campos.' });
      return;
    }

    this.apiService.login(this.authUsername, this.authPassword).subscribe({
      next: (res) => {
        this.apiService.saveToken(res.token);
        this.authSuccess.emit(); // Notificamos al componente padre
      },
      error: (err) => {
        const errorMsg = err.error?.error || 'Usuario o contraseña incorrectos.';
        Swal.fire({ icon: 'error', title: 'Error de acceso', text: errorMsg });
      }
    });
  }

  /**
   * Ejecuta el flujo de registro.
   */
  onRegister(): void {
    if (!this.authUsername || !this.authPassword) {
      Swal.fire({ icon: 'warning', title: 'Campos vacíos', text: 'Por favor, completa todos los campos.' });
      return;
    }

    this.apiService.register(this.authUsername, this.authPassword).subscribe({
      next: () => {
        Swal.fire({
          icon: 'success',
          title: '¡Registro Exitoso!',
          text: 'Ahora puedes iniciar sesión con tus credenciales.',
          confirmButtonText: 'Ir a Iniciar Sesión'
        }).then(() => {
          this.setView('login');
        });
      },
      error: (err) => {
        const errorMsg = err.error?.error || 'Ocurrió un error al registrar el usuario.';
        Swal.fire({ icon: 'error', title: 'Error de registro', text: errorMsg });
      }
    });
  }
}
