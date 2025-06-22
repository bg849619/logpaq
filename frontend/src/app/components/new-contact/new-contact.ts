import { Component } from '@angular/core';
import { FormsModule } from '@angular/forms';

@Component({
  selector: 'app-new-contact',
  imports: [FormsModule],
  templateUrl: './new-contact.html',
  styleUrl: './new-contact.css'
})
export class NewContact {
  public callsign: string = '';
  public class: string = '';
  public section: string = '';

  constructor() {
    // Initialize any necessary properties or services here
  }

  // Method to handle form submission
  submitContact() {
    // Logic to handle the new contact submission
    console.log('New contact submitted:', {
      callsign: this.callsign,
      class: this.class,
      section: this.section

    });
  }
}
