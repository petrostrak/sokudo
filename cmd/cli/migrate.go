package main

func doMigrate(arg2, arg3 string) error {
	checkForDB()
	tx, err := skd.PopConnect()
	if err != nil {
		exitGracefully(err)
	}
	defer tx.Close()

	// run the migration command
	switch arg2 {
	case "up":
		err := skd.PopMigrateUp(tx)
		if err != nil {
			return err
		}
	case "down":
		if arg3 == "all" {
			err := skd.PopMigrateDown(tx, -1)
			if err != nil {
				return err
			}
		} else {
			err := skd.PopMigrateDown(tx, 1)
			if err != nil {
				return err
			}
		}
	case "reset":
		err := skd.PopMigrateReset(tx)
		if err != nil {
			return err
		}
	default:
		showHelp()
	}

	return nil
}
