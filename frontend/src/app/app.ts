import { Component, OnInit, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ApiService } from './core/services/api.service';
import { AuthComponent } from './features/auth/auth.component';
import { DashboardComponent } from './features/dashboard/dashboard.component';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [CommonModule, AuthComponent, DashboardComponent],
  templateUrl: './app.html',
  styleUrl: './app.css'
})
export class App implements OnInit {
  private apiService = inject(ApiService);

  // Determina si el usuario está autenticado para mostrar el dashboard o el login
  isAuthenticated = false;

  ngOnInit(): void {
    // Verificamos el estado de autenticación inicial
    this.isAuthenticated = this.apiService.isAuthenticated();
  }

  /**
   * Se ejecuta cuando el AuthComponent emite el evento de éxito.
   */
  onAuthSuccess(): void {
    this.isAuthenticated = true;
  }

  /**
   * Se ejecuta cuando el DashboardComponent emite el evento de cierre de sesión.
   */
  onLogout(): void {
    this.isAuthenticated = false;
  }
}
