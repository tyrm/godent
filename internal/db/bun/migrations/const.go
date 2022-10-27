package migrations

const (
	CreatingColumn    = "creating column '%s' in table `%s`"
	CreatingColumnErr = "can't create column '%s' in table `%s`: %s"
	DroppingColumn    = "dropping column '%s' in table `%s`"
	DroppingColumnErr = "can't drop column '%s' in table `%s`: %s"
	CreatingTable     = "creating table '%s'"
	CreatingTableErr  = "can't create table '%s': %s"
	DroppingTable     = "dropping table '%s'"
	DroppingTableErr  = "can't drop table '%s': %s"
)
