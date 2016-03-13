# test.sh takes a list of packages, appends them to a whole string and sends
# them to go test.
#
# Since using 'go test ./...' will take the vendor folder as a valid package
# path, it will attempt to test all of the vendored packages. Also, since
# this project contains a frontend folder with all of it's frontend
# npm/bower modules, 'go test ./...' will also scan those folders taking up
# more time for the test run to complete.
#
# Note: If a new package is added with tests, it has to be added to the
#       PACKAGE_LIST array.


PACKAGE_LIST=(
	"github.com/ab22/abcd"
	"github.com/ab22/abcd/config"
	"github.com/ab22/abcd/handlers"
	"github.com/ab22/abcd/handlers/auth"
	"github.com/ab22/abcd/handlers/static"
	"github.com/ab22/abcd/handlers/student"
	"github.com/ab22/abcd/handlers/user"
	"github.com/ab22/abcd/httputils"
	"github.com/ab22/abcd/models"
	"github.com/ab22/abcd/routes"
	"github.com/ab22/abcd/services"
)

ALL_PACKAGES=$(printf " %s" "${PACKAGE_LIST[@]}")

go test -v ${ALL_PACKAGES}

