package main

func doMigrate(arg2, arg3 string) error {
	dsn := getDSN()

	// run the migration command
	switch arg2 {
	case "up":
		err := skd.MigrateUp(dsn)
		if err != nil {
			return err
		}

	case "down":
		if arg3 == "all" {
			err := skd.MigrateDownAll(dsn)
			if err != nil {
				return err
			}
		} else {
			err := skd.Steps(-1, dsn)
			if err != nil {
				return err
			}
		}
	case "reset":
		err := skd.MigrateDownAll(dsn)
		if err != nil {
			return err
		}
		err = skd.MigrateUp(dsn)
		if err != nil {
			return err
		}
	default:
		showHelp()
	}

	return nil
}
