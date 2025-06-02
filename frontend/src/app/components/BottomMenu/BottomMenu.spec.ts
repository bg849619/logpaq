import { ComponentFixture, TestBed } from '@angular/core/testing';

import { BottomMenu } from './BottomMenu';

describe('BottomMenu', () => {
  let component: BottomMenu;
  let fixture: ComponentFixture<BottomMenu>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [BottomMenu]
    })
    .compileComponents();

    fixture = TestBed.createComponent(BottomMenu);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
