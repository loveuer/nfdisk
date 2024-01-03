import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DelItemComponent } from './del-item.component';

describe('DelItemComponent', () => {
  let component: DelItemComponent;
  let fixture: ComponentFixture<DelItemComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [DelItemComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(DelItemComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
