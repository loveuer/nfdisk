import { ComponentFixture, TestBed } from '@angular/core/testing';

import { NewBucketComponent } from './new-bucket.component';

describe('NewBucketComponent', () => {
  let component: NewBucketComponent;
  let fixture: ComponentFixture<NewBucketComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [NewBucketComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(NewBucketComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
