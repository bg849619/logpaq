import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { BottomMenu } from './components/BottomMenu/BottomMenu';
import { ContactList } from './components/contact-list/contact-list';

@Component({
  selector: 'app-root',
  imports: [RouterOutlet, BottomMenu, ContactList],
  templateUrl: './app.html',
  styleUrl: './app.css'
})
export class App {
  protected title = 'LogPAQ';
}
