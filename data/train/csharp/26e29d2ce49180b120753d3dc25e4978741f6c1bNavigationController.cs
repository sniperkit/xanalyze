// This file has been autogenerated from a class added in the UI designer.
using System;
using Foundation;
using UIKit;
using TransitionsDemo.InteractionControllers;
using ObjCRuntime;
using System.Runtime.CompilerServices;

namespace TransitionsDemo
{
	public partial class NavigationController : UINavigationController
	{
		[Export ("initWithCoder:")]
		public NavigationController (NSCoder coder) : base (coder)
		{
			WeakDelegate = this;
		}

		[Export ("navigationController:animationControllerForOperation:fromViewController:toViewController:")]
		public IUIViewControllerAnimatedTransitioning GetAnimationControllerForOperation (UINavigationController navigationController,
		                                                                                  UINavigationControllerOperation operation,
		                                                                                  UIViewController fromViewController,
		                                                                                  UIViewController toViewController)
		{
			if (AppDelegate.NavigationControllerInteractionController != null) {
				AppDelegate.NavigationControllerInteractionController.WireToViewController (toViewController,
				                                                                            CEInteractionOperation.Pop);
			}

			if (AppDelegate.NavigationControllerAnimationController != null) {
				bool reverse = operation == UINavigationControllerOperation.Pop;
				AppDelegate.NavigationControllerAnimationController.Reverse = reverse;
			}

			return AppDelegate.NavigationControllerAnimationController;
		}
	}
}
