//
//  MobileEbitenViewControllerWithErrorHandling.m
//  shaderuniformtest
//
//  Created by ichibankunio on 2022/10/14.
//

#import "MobileEbitenViewControllerWithErrorHandling.h"

#import <Foundation/Foundation.h>

@implementation MobileEbitenViewControllerWithErrorHandling {
}

- (void)onErrorOnGameUpdate:(NSError*)err {
    // You can define your own error handling e.g., using Crashlytics.
    NSLog(@"Inovation Error!: %@", err);
}

@end
