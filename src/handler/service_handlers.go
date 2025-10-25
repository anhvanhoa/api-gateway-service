package handler

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	proto_auth "github.com/anhvanhoa/sf-proto/gen/auth/v1"
	proto_user "github.com/anhvanhoa/sf-proto/gen/user/v1"
	proto_session "github.com/anhvanhoa/sf-proto/gen/session/v1"
	proto_role "github.com/anhvanhoa/sf-proto/gen/role/v1"
	proto_permission "github.com/anhvanhoa/sf-proto/gen/permission/v1"
	proto_role_permission "github.com/anhvanhoa/sf-proto/gen/role_permission/v1"
	proto_user_role "github.com/anhvanhoa/sf-proto/gen/user_role/v1"
	proto_resource_permission "github.com/anhvanhoa/sf-proto/gen/resource_permission/v1"
	proto_greenhouse "github.com/anhvanhoa/sf-proto/gen/greenhouse/v1"
	proto_greenhouse_installation_log "github.com/anhvanhoa/sf-proto/gen/greenhouse_installation_log/v1"
	proto_growing_zone "github.com/anhvanhoa/sf-proto/gen/growing_zone/v1"
	proto_growing_zone_history "github.com/anhvanhoa/sf-proto/gen/growing_zone_history/v1"
	proto_iot_device "github.com/anhvanhoa/sf-proto/gen/iot_device/v1"
	proto_iot_device_history "github.com/anhvanhoa/sf-proto/gen/iot_device_history/v1"
	proto_device_type "github.com/anhvanhoa/sf-proto/gen/device_type/v1"
	proto_sensor_data "github.com/anhvanhoa/sf-proto/gen/sensor_data/v1"
	proto_plant_variety "github.com/anhvanhoa/sf-proto/gen/plant_variety/v1"
	proto_planting_cycle "github.com/anhvanhoa/sf-proto/gen/planting_cycle/v1"
	proto_harvest_record "github.com/anhvanhoa/sf-proto/gen/harvest_record/v1"
	proto_pest_disease_record "github.com/anhvanhoa/sf-proto/gen/pest_disease_record/v1"
	proto_irrigation_schedule "github.com/anhvanhoa/sf-proto/gen/irrigation_schedule/v1"
	proto_irrigation_log "github.com/anhvanhoa/sf-proto/gen/irrigation_log/v1"
	proto_fertilizer_type "github.com/anhvanhoa/sf-proto/gen/fertilizer_type/v1"
	proto_fertilizer_schedule "github.com/anhvanhoa/sf-proto/gen/fertilizer_schedule/v1"
	proto_maintenance_schedule "github.com/anhvanhoa/sf-proto/gen/maintenance_schedule/v1"
	proto_environmental_alert "github.com/anhvanhoa/sf-proto/gen/environmental_alert/v1"
	proto_status_history "github.com/anhvanhoa/sf-proto/gen/status_history/v1"
	proto_mail_provider "github.com/anhvanhoa/sf-proto/gen/mail_provider/v1"
	proto_mail_status "github.com/anhvanhoa/sf-proto/gen/mail_status/v1"
	proto_mail_history "github.com/anhvanhoa/sf-proto/gen/mail_history/v1"
	proto_mail_tmpl "github.com/anhvanhoa/sf-proto/gen/mail_tmpl/v1"
	proto_type_mail "github.com/anhvanhoa/sf-proto/gen/type_mail/v1"
	proto_system_configuration "github.com/anhvanhoa/sf-proto/gen/system_configuration/v1"
	proto_cost_tracking "github.com/anhvanhoa/sf-proto/gen/cost_tracking/v1"
	proto_media "github.com/anhvanhoa/sf-proto/gen/media/v1"
)

type ServiceHandler struct {
	Handler func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error
	Swagger string
}

// GetServiceHandlers trả về map chứa tất cả service handlers
func GetServiceHandlers() map[string]ServiceHandler {
	return map[string]ServiceHandler{
		"auth": {
			Handler: proto_auth.RegisterAuthServiceHandler,
			Swagger: "./swagger/auth/v1/auth.swagger.json",
		},
		"user": {
			Handler: proto_user.RegisterUserServiceHandler,
			Swagger: "./swagger/user/v1/user.swagger.json",
		},
		"session": {
			Handler: proto_session.RegisterSessionServiceHandler,
			Swagger: "./swagger/session/v1/session.swagger.json",
		},
		"role": {
			Handler: proto_role.RegisterRoleServiceHandler,
			Swagger: "./swagger/role/v1/role.swagger.json",
		},
		"permission": {
			Handler: proto_permission.RegisterPermissionServiceHandler,
			Swagger: "./swagger/permission/v1/permission.swagger.json",
		},
		"role_permission": {
			Handler: proto_role_permission.RegisterRolePermissionServiceHandler,
			Swagger: "./swagger/role_permission/v1/role_permission.swagger.json",
		},
		"user_role": {
			Handler: proto_user_role.RegisterUserRoleServiceHandler,
			Swagger: "./swagger/user_role/v1/user_role.swagger.json",
		},
		"resource_permission": {
			Handler: proto_resource_permission.RegisterResourcePermissionServiceHandler,
			Swagger: "./swagger/resource_permission/v1/resource_permission.swagger.json",
		},
		"greenhouse": {
			Handler: proto_greenhouse.RegisterGreenhouseServiceHandler,
			Swagger: "./swagger/greenhouse/v1/greenhouse.swagger.json",
		},
		"greenhouse_installation_log": {
			Handler: proto_greenhouse_installation_log.RegisterGreenhouseInstallationLogServiceHandler,
			Swagger: "./swagger/greenhouse_installation_log/v1/greenhouse_installation_log.swagger.json",
		},
		"growing_zone": {
			Handler: proto_growing_zone.RegisterGrowingZoneServiceHandler,
			Swagger: "./swagger/growing_zone/v1/growing_zone.swagger.json",
		},
		"growing_zone_history": {
			Handler: proto_growing_zone_history.RegisterGrowingZoneHistoryServiceHandler,
			Swagger: "./swagger/growing_zone_history/v1/growing_zone_history.swagger.json",
		},
		"iot_device": {
			Handler: proto_iot_device.RegisterIoTDeviceServiceHandler,
			Swagger: "./swagger/iot_device/v1/iot_device.swagger.json",
		},
		"iot_device_history": {
			Handler: proto_iot_device_history.RegisterIoTDeviceHistoryServiceHandler,
			Swagger: "./swagger/iot_device_history/v1/iot_device_history.swagger.json",
		},
		"device_type": {
			Handler: proto_device_type.RegisterDeviceTypeServiceHandler,
			Swagger: "./swagger/device_type/v1/device_type.swagger.json",
		},
		"sensor_data": {
			Handler: proto_sensor_data.RegisterSensorDataServiceHandler,
			Swagger: "./swagger/sensor_data/v1/sensor_data.swagger.json",
		},
		"plant_variety": {
			Handler: proto_plant_variety.RegisterPlantVarietyServiceHandler,
			Swagger: "./swagger/plant_variety/v1/plant_variety.swagger.json",
		},
		"planting_cycle": {
			Handler: proto_planting_cycle.RegisterPlantingCycleServiceHandler,
			Swagger: "./swagger/planting_cycle/v1/planting_cycle.swagger.json",
		},
		"harvest_record": {
			Handler: proto_harvest_record.RegisterHarvestRecordServiceHandler,
			Swagger: "./swagger/harvest_record/v1/harvest_record.swagger.json",
		},
		"pest_disease_record": {
			Handler: proto_pest_disease_record.RegisterPestDiseaseRecordServiceHandler,
			Swagger: "./swagger/pest_disease_record/v1/pest_disease_record.swagger.json",
		},
		"irrigation_schedule": {
			Handler: proto_irrigation_schedule.RegisterIrrigationScheduleServiceHandler,
			Swagger: "./swagger/irrigation_schedule/v1/irrigation_schedule.swagger.json",
		},
		"irrigation_log": {
			Handler: proto_irrigation_log.RegisterIrrigationLogServiceHandler,
			Swagger: "./swagger/irrigation_log/v1/irrigation_log.swagger.json",
		},
		"fertilizer_type": {
			Handler: proto_fertilizer_type.RegisterFertilizerTypeServiceHandler,
			Swagger: "./swagger/fertilizer_type/v1/fertilizer_type.swagger.json",
		},
		"fertilizer_schedule": {
			Handler: proto_fertilizer_schedule.RegisterFertilizerScheduleServiceHandler,
			Swagger: "./swagger/fertilizer_schedule/v1/fertilizer_schedule.swagger.json",
		},
		"maintenance_schedule": {
			Handler: proto_maintenance_schedule.RegisterMaintenanceScheduleServiceHandler,
			Swagger: "./swagger/maintenance_schedule/v1/maintenance_schedule.swagger.json",
		},
		"environmental_alert": {
			Handler: proto_environmental_alert.RegisterEnvironmentalAlertServiceHandler,
			Swagger: "./swagger/environmental_alert/v1/environmental_alert.swagger.json",
		},
		"status_history": {
			Handler: proto_status_history.RegisterStatusHistoryServiceHandler,
			Swagger: "./swagger/status_history/v1/status_history.swagger.json",
		},
		"mail_provider": {
			Handler: proto_mail_provider.RegisterMailProviderServiceHandler,
			Swagger: "./swagger/mail_provider/v1/mail_provider.swagger.json",
		},
		"mail_status": {
			Handler: proto_mail_status.RegisterMailStatusServiceHandler,
			Swagger: "./swagger/mail_status/v1/mail_status.swagger.json",
		},
		"mail_history": {
			Handler: proto_mail_history.RegisterMailHistoryServiceHandler,
			Swagger: "./swagger/mail_history/v1/mail_history.swagger.json",
		},
		"mail_tmpl": {
			Handler: proto_mail_tmpl.RegisterMailTmplServiceHandler,
			Swagger: "./swagger/mail_tmpl/v1/mail_tmpl.swagger.json",
		},
		"type_mail": {
			Handler: proto_type_mail.RegisterTypeMailServiceHandler,
			Swagger: "./swagger/type_mail/v1/type_mail.swagger.json",
		},
		"system_configuration": {
			Handler: proto_system_configuration.RegisterSystemConfigurationServiceHandler,
			Swagger: "./swagger/system_configuration/v1/system_configuration.swagger.json",
		},
		"cost_tracking": {
			Handler: proto_cost_tracking.RegisterCostTrackingServiceHandler,
			Swagger: "./swagger/cost_tracking/v1/cost_tracking.swagger.json",
		},
		"media": {
			Handler: proto_media.RegisterMediaServiceHandler,
			Swagger: "./swagger/media/v1/media.swagger.json",
		},
	}
}
