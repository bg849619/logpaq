import { Component, Inject } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { MatDialogModule, MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';

type OperatorMenuData = {
  callsign: string;
  name: string;
}

@Component({
  selector: 'app-operator-menu',
  imports: [MatDialogModule, FormsModule],
  templateUrl: './operator-menu.html',
  styleUrl: './operator-menu.css'
})
export class OperatorMenu {
  constructor(
    public dialogRef: MatDialogRef<OperatorMenu>, @Inject(MAT_DIALOG_DATA) public data: OperatorMenuData
  ) {}

  onNoClick(): void {
    this.dialogRef.close();
  }
}
