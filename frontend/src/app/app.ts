import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { BottomMenu } from './components/BottomMenu/BottomMenu';

@Component({
  selector: 'app-root',
  imports: [RouterOutlet, BottomMenu],
  templateUrl: './app.html',
  styleUrl: './app.css'
})
export class App {
  protected title = 'dumbass';
}
