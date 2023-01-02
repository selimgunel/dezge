import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule, Routes } from '@angular/router';
import { TournamentListComponent } from './components/tournament-list/tournament-list.component';
import { TournamentViewComponent } from './components/tournament-view/tournament-view.component';

const routes: Routes = [
  {path: '', component: TournamentListComponent},
]

@NgModule({
  declarations: [
    TournamentListComponent,
    TournamentViewComponent
  ],
  imports: [
    CommonModule,
    RouterModule.forChild(routes)
  ]
})
export class TournamentModule { }
