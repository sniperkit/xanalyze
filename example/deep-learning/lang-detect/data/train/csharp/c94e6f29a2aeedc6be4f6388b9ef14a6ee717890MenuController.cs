// This file has been autogenerated from a class added in the UI designer.

using Foundation;
using System;
using System.CodeDom.Compiler;
using UIKit;
using ConservappiOsAndroid.Model;
using System.Threading.Tasks;

namespace ConservappiOsAndroid.iOS{
	public partial class MenuController : BaseController{
		UIAlertView alert = new UIAlertView();


		Rootobject sesion;

		public MenuController(IntPtr handle) : base(handle) {
		}

		//public async override void ViewWillAppear(bool animated) {
		//	base.ViewWillAppear(animated);
		//	if (await WaitDataUserAsync() == true) {
		//		nombreMenuLabel.Text = sesion.sesion.nombre_completo;

		//	}
		//}
		public  override void ViewDidLoad() {
			base.ViewDidLoad();

			//var contentController = (ContentController)Storyboard.InstantiateViewController("ContentController");
			//var otherController = (OtherController)Storyboard.InstantiateViewController("OtherController");
			//var introController = (IntroController)Storyboard.InstantiateViewController("IntroController");
			var estadoCuentaController = (EstadoDeCuentaTableViewController)Storyboard.InstantiateViewController("EstadoDeCuentaTableViewController");
			var historialEstadoCuenta = (HistorialEstadoCuentaTableViewController)Storyboard.InstantiateViewController("HistorialEstadoCuentaTableViewController");

			var payReference = (PayRefrenceTableViewController)Storyboard.InstantiateViewController("PayRefrenceTableViewController");
			var amortizacion = (AmortizacionTableViewController)Storyboard.InstantiateViewController("AmortizacionTableViewController");

			var movimientos = (MovimientosTableViewController)Storyboard.InstantiateViewController("MovimientosTableViewController");

			var notify= (SwiperTableViewController)Storyboard.InstantiateViewController("SwiperTableViewController");

			var garantias = (GarantiasTableViewController)Storyboard.InstantiateViewController("GarantiasTableViewController");
			//nombreMenuLabel.Text = sesion.sesion.nombre_completo;

			//ContentButton.TouchUpInside += (o, e) => {

			//	if (NavController.TopViewController != estadoCuentaController) {
			//		NavController.PopViewController(true);
			//		//SidebarController.CloseMenu();				
			//	}
			//	Console.WriteLine(NavController.TopViewController.ToString());

			//	if (NavController.TopViewController as ContentController == null) {
			//		Console.WriteLine("Content Controller == null");
			//		Console.WriteLine(NavController.TopViewController.ToString());


			//		NavController.PushViewController(contentController, true);

			//	} else {
			//		SidebarController.CloseMenu();
			//	}

			//	SidebarController.CloseMenu();
			//};

			//OtherViewButton.TouchUpInside += (sender, e) => {

			//	if (NavController.TopViewController != estadoCuentaController) {
			//		NavController.PopViewController(true);
			//		//SidebarController.CloseMenu();				
			//	}
			//	Console.WriteLine(NavController.TopViewController.ToString());

			//	if (NavController.TopViewController as OtherController == null) {//if (NavController.TopViewController as OtherController == null) 
			//		Console.WriteLine("Other Controller == null");
			//		Console.WriteLine(NavController.TopViewController.ToString());


			//		NavController.PushViewController(otherController, true);

			//	} else {
			//		SidebarController.CloseMenu();
			//	}
			//	//if (NavController.TopViewController as OtherController == null) {
			//	//	//NavController.PushViewController(introController, false);

			//	//	NavController.PushViewController(otherController, false);


			//	//}


			//	SidebarController.CloseMenu();


			//};

			HistorialViewButton.TouchUpInside += (sender, e) => { 
				if (NavController.TopViewController != estadoCuentaController) {
					NavController.PopViewController(true);
					//SidebarController.CloseMenu();				
				}
				Console.WriteLine(NavController.TopViewController.ToString());

				if (NavController.TopViewController as HistorialEstadoCuentaTableViewController == null) {// (NavController.TopViewController as OtherController == null) 
					Console.WriteLine("Other Controller == null");
					Console.WriteLine(NavController.TopViewController.ToString());


					NavController.PushViewController(historialEstadoCuenta, true);

				} else {
					SidebarController.CloseMenu();
				}
				//if (NavController.TopViewController as OtherController == null) {
				//	//NavController.PushViewController(introController, false);

				//	NavController.PushViewController(otherController, false);


				//}


				SidebarController.CloseMenu();
				
			};

			ReferenciaPagoViewButton.TouchUpInside += (sender, e) => {
				if (NavController.TopViewController != estadoCuentaController) {
					NavController.PopViewController(true);
					//SidebarController.CloseMenu();				
				}
				Console.WriteLine(NavController.TopViewController.ToString());

				if (NavController.TopViewController as PayRefrenceTableViewController == null) {//NavController.TopViewController as OtherController == null (Sentencia Anterior Posible Error)
					Console.WriteLine("Other Controller == null");
					Console.WriteLine(NavController.TopViewController.ToString());


					NavController.PushViewController(payReference, true);

				} else {
					SidebarController.CloseMenu();
				}
				//if (NavController.TopViewController as OtherController == null) {
				//	//NavController.PushViewController(introController, false);

				//	NavController.PushViewController(otherController, false);


				//}


				SidebarController.CloseMenu();
				
			};

			AmortizacionViewButton.TouchUpInside += (sender, e) => { 
				if (NavController.TopViewController != estadoCuentaController) {
					NavController.PopViewController(true);
					//SidebarController.CloseMenu();				
				}
				Console.WriteLine(NavController.TopViewController.ToString());

				if (NavController.TopViewController as AmortizacionTableViewController == null) {
					Console.WriteLine("Other Controller == null");
					Console.WriteLine(NavController.TopViewController.ToString());


					NavController.PushViewController(amortizacion, true);

				} else {
					SidebarController.CloseMenu();
				}
				//if (NavController.TopViewController as OtherController == null) {
				//	//NavController.PushViewController(introController, false);

				//	NavController.PushViewController(otherController, false);


				//}


				SidebarController.CloseMenu();
				
			};

			estadoCuentaViewButton.TouchUpInside += (sender, e) => { 
				if (NavController.TopViewController != estadoCuentaController) {
					NavController.PopViewController(true);
					//SidebarController.CloseMenu();				
				}
				Console.WriteLine(NavController.TopViewController.ToString());

				if (NavController.TopViewController as EstadoDeCuentaTableViewController == null) {
					Console.WriteLine("Other Controller == null");
					Console.WriteLine(NavController.TopViewController.ToString());


					NavController.PushViewController(estadoCuentaController, true);

				} else {
					SidebarController.CloseMenu();
				}
				//if (NavController.TopViewController as OtherController == null) {
				//	//NavController.PushViewController(introController, false);

				//	NavController.PushViewController(otherController, false);


				//}


				SidebarController.CloseMenu();
				
			};

			NotificationViewButton.TouchUpInside += (sender, e) => { 
				if (NavController.TopViewController != estadoCuentaController) {
					NavController.PopViewController(true);
					//SidebarController.CloseMenu();				
				}
				Console.WriteLine(NavController.TopViewController.ToString());

				if (NavController.TopViewController as SwiperTableViewController == null) {
					Console.WriteLine("Other Controller == null");
					Console.WriteLine(NavController.TopViewController.ToString());


					NavController.PushViewController(notify, true);

				} else {
					SidebarController.CloseMenu();
				}
				//if (NavController.TopViewController as OtherController == null) {
				//	//NavController.PushViewController(introController, false);

				//	NavController.PushViewController(otherController, false);


				//}


				SidebarController.CloseMenu();
				
			};

			garantiasViewButton.TouchUpInside += (sender, e) => { 
				if (NavController.TopViewController != estadoCuentaController) {
					NavController.PopViewController(true);
					//SidebarController.CloseMenu();				
				}
				Console.WriteLine(NavController.TopViewController.ToString());

				if (NavController.TopViewController as GarantiasTableViewController == null) {
					Console.WriteLine("GarantiasTableViewController == null");
					Console.WriteLine(NavController.TopViewController.ToString());


					NavController.PushViewController(garantias, true);

				} else {
					SidebarController.CloseMenu();
				}
				//if (NavController.TopViewController as OtherController == null) {
				//	//NavController.PushViewController(introController, false);

				//	NavController.PushViewController(otherController, false);


				//}


				SidebarController.CloseMenu();
			};



			closeSessionButton.TouchUpInside += (sender, e) => { 
				var appDelegate = UIApplication.SharedApplication.Delegate as AppDelegate;

				var rootController = new RootViewController();

				var loginView = (PrincipalNavViewController)Storyboard.InstantiateViewController("PrincipalNavViewController");

				appDelegate.SetRootViewController(loginView, true);


			};

		}



		public async override void ViewWillAppear(bool animated) {
			base.ViewWillAppear(animated);
			if (await WaitDataUserAsync() == true) {
				Console.WriteLine("Ready data session!");
			}
		}


		async Task<bool> WaitDataUserAsync() {
			Console.WriteLine("WaitHistorialEstadoCuenta Method...");
			try {
				Service.Service serv = new Service.Service();
				REstado rEstado = await serv.EstadoCuentaAsync("1", "0", CurrentSession.identificador);


				if (rEstado.codigo == 1000) {

					Console.WriteLine("Session data ready!");
					Console.WriteLine("Nombre De cliente: "+ CurrentSession.nombre_completo);
					nombreSession.Text = CurrentSession.nombre_completo;
					numeroSession.Text = CurrentSession.no_telefono;

					return true;
				} else { 
					alert.Title = "Error de Red";
					alert.Message = "Verifique su conexión a internet";
					alert.AddButton("OK");
					alert.Show();
					Console.WriteLine("Session data error!");

					return false;
				}


			}
			catch { 
				alert.Title = "Error de Red";
				alert.Message = "Verifique su conexión a internet";
				alert.AddButton("OK");
				alert.Show();
				Console.WriteLine("Session data error!");

				
			}
			return false;
		
		}
	}
}
