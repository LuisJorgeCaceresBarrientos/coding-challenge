import { Injectable, inject } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ApiService {
  private http = inject(HttpClient);
  // URL base de la API en Go (detecta automáticamente si es localhost o la IP del servidor)
  private baseURL = `http://${window.location.hostname}:8080/api`;

  /**
   * Guarda el token JWT en el localStorage del navegador.
   * @param token El token a guardar.
   */
  saveToken(token: string): void {
    localStorage.setItem('auth_token', token);
  }

  /**
   * Obtiene el token JWT almacenado.
   * @returns El token o null si no existe.
   */
  getToken(): string | null {
    return localStorage.getItem('auth_token');
  }

  /**
   * Elimina el token JWT de la sesión (Cierre de sesión).
   */
  logout(): void {
    localStorage.removeItem('auth_token');
  }

  /**
   * Verifica si el usuario actual tiene una sesión válida (tiene token).
   * @returns Verdadero si está autenticado.
   */
  isAuthenticated(): boolean {
    return this.getToken() !== null;
  }

  /**
   * Llama al endpoint de registro.
   */
  register(username: string, password: string): Observable<any> {
    return this.http.post(`${this.baseURL}/auth/register`, { username, password });
  }

  /**
   * Llama al endpoint de inicio de sesión.
   */
  login(username: string, password: string): Observable<{ token: string }> {
    return this.http.post<{ token: string }>(`${this.baseURL}/auth/login`, { username, password });
  }

  /**
   * Envía la matriz a la API de Go para su rotación y factorización QR.
   * Inyecta el token JWT en la cabecera de la petición.
   */
  processMatrix(matrix: number[][]): Observable<any> {
    const token = this.getToken();
    const headers = new HttpHeaders({
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${token}`
    });

    return this.http.post(`${this.baseURL}/matrix/process`, { matrix }, { headers });
  }
}
