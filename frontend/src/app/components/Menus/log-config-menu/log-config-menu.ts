import { Component, Inject } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { MatDialogModule, MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { LogConfig } from '../../../../types';

@Component({
  selector: 'app-log-config-menu',
  imports: [FormsModule, MatDialogModule],
  templateUrl: './log-config-menu.html',
  styleUrl: './log-config-menu.css'
})
export class LogConfigMenu {
  constructor(
    public dialogRef: MatDialogRef<LogConfigMenu>, @Inject(MAT_DIALOG_DATA) public data: LogConfig
  ) {}

  onNoClick(): void {
    this.dialogRef.close();
  }
}
