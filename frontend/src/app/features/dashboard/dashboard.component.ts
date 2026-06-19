import { Component, OnInit, Output, EventEmitter, inject, ChangeDetectorRef } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { ApiService } from '../../core/services/api.service';

@Component({
  selector: 'app-dashboard',
  standalone: true,
  imports: [CommonModule, FormsModule],
  templateUrl: './dashboard.component.html'
})
export class DashboardComponent implements OnInit {
  private apiService = inject(ApiService);
  private cdr = inject(ChangeDetectorRef);

  // Evento que notifica al contenedor principal cuando se cierra la sesión
  @Output() logoutEvent = new EventEmitter<void>();

  // Estado del Dashboard
  errorMessage = '';
  matrixRows = 3;
  matrixCols = 3;
  matrixInput: number[][] = [];
  
  processingResult: any = null;
  isProcessing = false;

  ngOnInit(): void {
    this.generateEmptyMatrix();
  }

  /**
   * Cierra la sesión y notifica al padre.
   */
  onLogout(): void {
    this.apiService.logout();
    this.processingResult = null;
    this.logoutEvent.emit();
  }

  /**
   * Inicializa una matriz vacía según las filas y columnas configuradas.
   */
  generateEmptyMatrix(): void {
    if (this.matrixRows < 1) this.matrixRows = 1;
    if (this.matrixCols < 1) this.matrixCols = 1;
    if (this.matrixRows > 10) this.matrixRows = 10;
    if (this.matrixCols > 10) this.matrixCols = 10;

    const newMatrix: number[][] = [];
    for (let i = 0; i < this.matrixRows; i++) {
      const row: number[] = [];
      for (let j = 0; j < this.matrixCols; j++) {
        row.push(0);
      }
      newMatrix.push(row);
    }
    this.matrixInput = newMatrix;
  }

  /**
   * Llena la matriz con valores numéricos aleatorios (entre -50 y 100).
   */
  fillWithRandomValues(): void {
    for (let i = 0; i < this.matrixRows; i++) {
      for (let j = 0; j < this.matrixCols; j++) {
        this.matrixInput[i][j] = Math.floor(Math.random() * 150) - 50;
      }
    }
  }

  /**
   * Envía la matriz al backend y procesa la respuesta.
   */
  submitMatrix(): void {
    this.errorMessage = '';
    this.processingResult = null;
    this.isProcessing = true;

    // Convertimos los valores a números enteros o flotantes de forma segura
    const cleanMatrix = this.matrixInput.map(row => 
      row.map(val => Number(val) || 0)
    );

    this.apiService.processMatrix(cleanMatrix).subscribe({
      next: (res) => {
        this.processingResult = res;
        this.isProcessing = false;
        this.cdr.detectChanges();
      },
      error: (err) => {
        this.errorMessage = err.error?.error || 'Error al procesar la matriz.';
        this.isProcessing = false;
        this.cdr.detectChanges();
      }
    });
  }

  /**
   * Helper crucial para rendimiento en listas. Rastrea las filas por su índice.
   */
  trackByRow(index: number): number {
    return index;
  }

  /**
   * Helper crucial para rendimiento en listas. Rastrea las columnas por su índice.
   */
  trackByCol(index: number): number {
    return index;
  }
}
