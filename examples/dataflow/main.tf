// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

// These variables would commonly be defined as environment variables or sourced in a .env file

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_id" {}
variable "region" {}

provider "oci" {
  region           = "${var.region}"
  tenancy_ocid     = "${var.tenancy_ocid}"
  user_ocid        = "${var.user_ocid}"
  fingerprint      = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
}

variable "application_display_name" {
  default = "tf_app"
}

variable "application_driver_shape" {
  default = "VM.Standard2.1"
}

variable "application_executor_shape" {
  default = "VM.Standard2.1"
}

variable "application_file_uri" {}

variable "application_language" {
  default = "PYTHON"
}

variable "application_num_executors" {
  default = 1
}

variable "application_spark_version" {
  default = "2.4"
}

variable "invoke_run_display_name" {
  default = "tf_run"
}

resource "oci_dataflow_application" "tf_application" {
  #Required
  compartment_id = "${var.compartment_id}"
  display_name   = "${var.application_display_name}"
  driver_shape   = "${var.application_driver_shape}"
  executor_shape = "${var.application_executor_shape}"
  file_uri       = "${var.application_file_uri}"
  language       = "${var.application_language}"
  num_executors  = "${var.application_num_executors}"
  spark_version  = "${var.application_spark_version}"

  #Optional
  #arguments       = "${var.application_arguments}"
  #class_name      = "${var.application_class_name}"
  #configuration   = "${var.application_configuration}"
  #defined_tags    = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.application_defined_tags_value}")}"
  #description     = "${var.application_description}"
  #freeform_tags   = "${var.application_freeform_tags}"
  #logs_bucket_uri = "${var.application_logs_bucket_uri}"

  #parameters {
  #Required
  #name  = "${var.application_parameters_name}"
  #value = "${var.application_parameters_value}"
  #}

  #warehouse_bucket_uri = "${var.application_warehouse_bucket_uri}"
}

data "oci_dataflow_applications" "tf_applications" {
  #Required
  compartment_id = "${var.compartment_id}"

  #Optional
  display_name = "${var.application_display_name}"
}

resource "oci_dataflow_invoke_run" "tf_invoke_run" {
  #Required
  application_id = "${oci_dataflow_application.tf_application.id}"
  compartment_id = "${var.compartment_id}"
  display_name   = "${var.invoke_run_display_name}"

  #Optional
  #arguments       = "${var.invoke_run_arguments}"
  #configuration   = "${var.invoke_run_configuration}"
  #defined_tags    = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.invoke_run_defined_tags_value}")}"
  #driver_shape    = "${var.invoke_run_driver_shape}"
  #executor_shape  = "${var.invoke_run_executor_shape}"
  #freeform_tags   = "${var.invoke_run_freeform_tags}"
  #logs_bucket_uri = "${var.invoke_run_logs_bucket_uri}"
  #num_executors   = "${var.invoke_run_num_executors}"

  #parameters {
  #Required
  #name  = "${var.invoke_run_parameters_name}"
  #value = "${var.invoke_run_parameters_value}"
  #}

  #warehouse_bucket_uri = "${var.invoke_run_warehouse_bucket_uri}"
}

data "oci_dataflow_invoke_runs" "tf_invoke_runs" {
  #Required
  compartment_id = "${var.compartment_id}"

  #Optional
  application_id = "${oci_dataflow_application.tf_application.id}"
}

data "oci_dataflow_run_logs" "tf_run_logs" {
  #Required
  run_id = "${oci_dataflow_invoke_run.tf_invoke_run.id}"
}
