import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { MainComponent } from './main/main.component';
import { StepsComponent } from './steps/steps.component';



const routes: Routes = [
  { path: '', component: MainComponent },
  { path: 'steps-chart', component: StepsComponent }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
