package repositories

import "github.com/muplat/muplat-backend/models"

func (db *Database) GetPostgresConfig(deploymentName, projectName string) (*models.PostgresConfig, error) {
	ac := &models.PostgresConfig{}
	err := db.Connection.
		Model(&models.PostgresConfig{}).
		Where("deployment_name = ?", deploymentName).
		Where("project_name = ?", projectName).
		Take(ac).Error
	if err != nil {
		return nil, err
	}
	return ac, nil
}

func (db *Database) SavePostgresConfig(
	deploymentName,
	projectName,
	diskSize,
	internalEndpoint,
	database,
	credentialsSecret string,
) error {
	pc := &models.PostgresConfig{
		DeploymentName:    deploymentName,
		ProjectName:       projectName,
		DiskSize:          diskSize,
		InternalEndpoint:  internalEndpoint,
		Database:          database,
		CredentialsSecret: credentialsSecret,
	}

	err := db.Connection.Create(pc).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) DeletePostgresConfig(
	deploymentName,
	projectName string,
) error {
	pc := &models.PostgresConfig{
		DeploymentName: deploymentName,
		ProjectName:    projectName,
	}
	err := db.Connection.Model(&models.PostgresConfig{}).Delete(pc).Error
	if err != nil {
		return err
	}
	return nil
}
