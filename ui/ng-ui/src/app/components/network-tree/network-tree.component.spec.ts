import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { NetworkTreeComponent } from './network-tree.component';

describe('NetworkTreeComponent', () => {
  let component: NetworkTreeComponent;
  let fixture: ComponentFixture<NetworkTreeComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ NetworkTreeComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(NetworkTreeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
