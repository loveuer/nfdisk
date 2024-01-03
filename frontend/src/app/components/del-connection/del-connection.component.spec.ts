import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DelConnectionComponent } from './del-connection.component';

describe('DelConnectionComponent', () => {
  let component: DelConnectionComponent;
  let fixture: ComponentFixture<DelConnectionComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [DelConnectionComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(DelConnectionComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
