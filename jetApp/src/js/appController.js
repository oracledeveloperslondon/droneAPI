/**
 * Copyright (c) 2014, 2018, Oracle and/or its affiliates.
 * The Universal Permissive License (UPL), Version 1.0
 */
/*
 * Your application specific code will go here
 */
define(['ojs/ojcore', 'knockout', 'ojs/ojknockout', 'ojs/ojmodule'],
  function(oj, ko) {
     function ControllerViewModel() {
       var self = this;

      self.appName = ko.observable("Drone Watcher");
      self.messagesPage = ko.observable("messages");
      self.animationPage = ko.observable("animation");

     };

     return new ControllerViewModel();
  }
);
