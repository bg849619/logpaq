import { Component } from '@angular/core';

@Component({
  selector: 'app-contact-list',
  imports: [],
  templateUrl: './contact-list.html',
  styleUrl: './contact-list.css'
})
export class ContactList {
  public contacts = Array(200).fill({
    id: '000000', // Would actually be a UUID probably.
    callsign: 'KE8PAQ',
    class: '1A',
    section: 'OH',
    time: '2025-06-13T12:29:21Z',
    band: '20M',
    mode: 'SSB',
    operator: 'N8BAP',
    op_name: 'PH'
  });
}
