import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { LayoutComponent } from './_layout/layout.component';
import { BuildComponent } from './build/build.component';
import { CommonModule } from '@angular/common';

import { ReactiveFormsModule } from '@angular/forms';
import { MatCardModule } from '@angular/material/card';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatSelectModule } from '@angular/material/select';
import { MatInputModule } from '@angular/material/input';

const routes: Routes = [
  {
    path: '',
    component: LayoutComponent,
    children: [
      {
        path: 'dashboard',
        loadChildren: () => import('./dashboard/dashboard.module').then((m) => m.DashboardModule)
      },
      {
        path: 'build',
        component: BuildComponent
      },
      {
        path: 'manage',
        loadChildren: () => import('./manage/manage.module').then((m) => m.ManageModule)
      },
      {
        path: 'monitor',
        loadChildren: () => import('./monitor/monitor.module').then((m) => m.MonitorModule)
      },
      {
        path: 'plan',
        loadChildren: () => import('./plan/plan.module').then((m) => m.PlanModule)
      },
      {
        path: 'test',
        loadChildren: () => import('./test/test.module').then((m) => m.TestModule)
      },
      {
        path: 'user-management',
        loadChildren: () => import('../modules/user-management/user-management.module').then((m) => m.UserManagementModule)
      },
      {
        path: '',
        redirectTo: 'monitor',
        pathMatch: 'full'
      },
      {
        path: '**',
        redirectTo: 'errors/404',
        pathMatch: 'full'
      }
    ]
  }
];

@NgModule({
  declarations: [],
  imports: [
    CommonModule,
    RouterModule.forChild(routes),
    MatCardModule,
    MatFormFieldModule,
    ReactiveFormsModule,
    MatSelectModule,
    MatInputModule
  ],
  exports: [RouterModule]
})
export class PagesRoutingModule {}
