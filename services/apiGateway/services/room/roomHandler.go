package roomHandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	roompb "github.com/sri2103/chat_me/protos/room"
	"github.com/sri2103/chat_me/services/apiGateway/config"
	"go.uber.org/zap"
)

type roomHandler struct {
	logger      *zap.Logger
	roomService roompb.RoomServiceClient
}

func New(cfg *config.Config) *roomHandler {
	return &roomHandler{
		logger:      cfg.Log,
		roomService: cfg.RoomClientService,
	}
}

type CreateUserReq struct {
	CreatedBy string   `json:"created_by"`
	RoomName  string   `json:"room_name"`
	Users     []string `json:"users"`
}

type CreateUserRes struct {
	RoomId  string `json:"room_id"`
	Success bool   `json:"success"`
}

// create room

func (r *roomHandler) CreateRoom(c echo.Context) error {
	var b CreateUserReq
	err := c.Bind(&b)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "body error")
	}

	if err := c.Validate(b); err != nil {
		return err
	}

	var cr roompb.CreateRoomRequest
	cr.RoomName = b.RoomName
	cr.CreatedBy = b.CreatedBy
	cr.UserIds = b.Users

	crRes, err := r.roomService.CreateRoom(c.Request().Context(), &cr)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	var res CreateUserRes
	res.RoomId = crRes.RoomId
	res.Success = crRes.Success
	return c.JSON(http.StatusOK, res)

}

type FetchRoomResponse struct {
	RoomId    string   `json:"room_id"`
	RoomName  string   `json:"room_name"`
	CreatedBy string   `json:"created_by"`
	Users     []string `json:"users"`
}

// fetch room
func (r *roomHandler) FetchRoom(c echo.Context) error {
	roomId := c.QueryParam("roomId")
	res, err := r.roomService.GetRoomDetails(c.Request().Context(), &roompb.GetRoomRequest{
		RoomId: roomId,
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "internal error happened")
	}

	rs := &FetchRoomResponse{
		RoomId:    res.RoomId,
		RoomName:  res.RoomName,
		CreatedBy: res.CreatedBy,
		Users:     res.GetUserIds(),
	}
	return c.JSON(http.StatusOK, rs)
}

// updateRoom

type UpdateReq struct {
	RoomId   string   `json:"room_id"`
	RoomName string   `json:"room_name"`
	Users    []string `json:"users"`
}

type UpdateRes struct {
	Success bool
}

func (r *roomHandler) UpdateRoom(c echo.Context) error {
	var b UpdateReq
	err := c.Bind(&b)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	var req roompb.UpdateRoomRequest
	req.RoomName = b.RoomName
	req.RoomId = b.RoomId
	req.UserIds = b.Users
	res, err := r.roomService.UpdateRoomSettings(c.Request().Context(), &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var d UpdateRes
	d.Success = res.GetSuccess()
	return c.JSON(http.StatusOK, d)
}
