package models

type (
	V4GetHousesByWorldAndTownResponse struct {
		Houses struct {
			GuildhallList []HouseInformation `json:"guildhall_list"`
			HouseList     []HouseInformation `json:"house_list"`
			Town          string             `json:"town"`
			World         string             `json:"world"`
		} `json:"houses"`
		Information APIInformation `json:"information"`
	}

	V4GetHouseByTownIDResponse struct {
		House       HouseExtendedInformation `json:"house"`
		Information APIInformation           `json:"information"`
	}

	HouseInformation struct {
		Auction   AuctionInformation `json:"auction"`
		Auctioned bool               `json:"auctioned"`
		HouseID   int                `json:"house_id"`
		Name      string             `json:"name"`
		Rent      int                `json:"rent"`
		Rented    bool               `json:"rented"`
		Size      int                `json:"size"`
	}

	AuctionInformation struct {
		CurrentBid int    `json:"current_bid"`
		Finished   bool   `json:"finished"`
		TimeLeft   string `json:"time_left"`
	}

	HouseExtendedInformation struct {
		Beds    int               `json:"beds"`
		Houseid int               `json:"houseid"`
		Img     string            `json:"img"`
		Name    string            `json:"name"`
		Rent    int               `json:"rent"`
		Size    int               `json:"size"`
		Status  StatusInformation `json:"status"`
		Town    string            `json:"town"`
		Type    string            `json:"type"`
		World   string            `json:"world"`
	}

	AuctionExtendedInfo struct {
		AuctionEnd     string `json:"auction_end"`
		AuctionOngoing bool   `json:"auction_ongoing"`
		CurrentBid     int    `json:"current_bid"`
		CurrentBidder  string `json:"current_bidder"`
	}

	RentalInformation struct {
		MovingDate       string `json:"moving_date"`
		Owner            string `json:"owner"`
		OwnerSex         string `json:"owner_sex"`
		PaidUntil        string `json:"paid_until"`
		TransferAccept   bool   `json:"transfer_accept"`
		TransferPrice    int    `json:"transfer_price"`
		TransferReceiver string `json:"transfer_receiver"`
	}

	StatusInformation struct {
		Auction       AuctionExtendedInfo `json:"auction"`
		IsAuctioned   bool                `json:"is_auctioned"`
		IsMoving      bool                `json:"is_moving"`
		IsRented      bool                `json:"is_rented"`
		IsTransfering bool                `json:"is_transfering"`
		Original      string              `json:"original"`
		Rental        RentalInformation   `json:"rental"`
	}
)
