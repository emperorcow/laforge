import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { NetworkComponent } from './network/network.component';
import { HostComponent } from './host/host.component';
import { HostTreeComponent } from "./host-tree/host-tree.component";
import { TeamComponent } from './team/team.component';
import { HostModalComponent } from './host-modal/host-modal.component';
import { MatDialogModule } from '@angular/material/dialog';
import { MatTableModule } from '@angular/material/table';
import { MatButtonModule } from '@angular/material/button';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { NetworkModalComponent } from './network-modal/network-modal.component';
import { StepComponent } from './step/step.component';
import { MomentModule } from 'ngx-moment';
import { LaforgePipesModule } from '../pipes/pipes.module';
import { TeamTreeComponent } from './team-tree/team-tree.component';
import { NetworkTreeComponent } from './network-tree/network-tree.component';

@NgModule({
  declarations: [NetworkComponent, HostComponent, TeamComponent, HostModalComponent, NetworkModalComponent, StepComponent, HostTreeComponent, TeamTreeComponent, NetworkTreeComponent],
  imports: [CommonModule, MatDialogModule, MatTableModule, MatButtonModule, MatCheckboxModule, MomentModule, LaforgePipesModule],
  exports: [NetworkComponent, HostComponent, TeamComponent, TeamTreeComponent, NetworkTreeComponent, HostTreeComponent]
})
export class ViewComponentsModule {}
