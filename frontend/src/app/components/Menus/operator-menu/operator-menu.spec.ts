import { ComponentFixture, TestBed } from '@angular/core/testing';

import { OperatorMenu } from './operator-menu';

describe('OperatorMenu', () => {
  let component: OperatorMenu;
  let fixture: ComponentFixture<OperatorMenu>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [OperatorMenu]
    })
    .compileComponents();

    fixture = TestBed.createComponent(OperatorMenu);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
