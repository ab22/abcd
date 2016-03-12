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

go test ${ALL_PACKAGES}

