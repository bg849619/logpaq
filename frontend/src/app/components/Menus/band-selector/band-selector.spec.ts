import { ComponentFixture, TestBed } from '@angular/core/testing';

import { BandSelector } from './band-selector';

describe('BandSelector', () => {
  let component: BandSelector;
  let fixture: ComponentFixture<BandSelector>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [BandSelector]
    })
    .compileComponents();

    fixture = TestBed.createComponent(BandSelector);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
