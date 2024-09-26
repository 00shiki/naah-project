package vouchers

import (
	"api-gateway/entity/vouchers"
	pb "api-gateway/proto"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

type VoucherService struct {
	client pb.VoucherServiceClient
}

func NewVoucherService(client pb.VoucherServiceClient) *VoucherService {
	return &VoucherService{
		client: client,
	}
}

func (vs *VoucherService) CreateVoucher(voucher *vouchers.Voucher) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req := &pb.AddVoucherRequest{
		VoucherId:  voucher.VoucherID,
		Discount:   voucher.Discount,
		ValidUntil: voucher.ValidUntil,
	}
	_, err := vs.client.AddVoucher(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func (vs *VoucherService) GetVoucherByID(voucherID string) (*vouchers.Voucher, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req := &pb.GetVoucherRequest{
		VoucherId: voucherID,
	}
	res, err := vs.client.GetVoucher(ctx, req)
	if err != nil {
		return nil, err
	}
	voucher := &vouchers.Voucher{
		VoucherID:  voucherID,
		Discount:   res.Discount,
		ValidUntil: res.ValidUntil,
		Used:       res.Used,
	}
	return voucher, nil
}

func (vs *VoucherService) GetVouchers() ([]vouchers.Voucher, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := vs.client.GetVoucherList(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	vouchersList := make([]vouchers.Voucher, len(res.Vouchers))
	for i, v := range res.Vouchers {
		vouchersList[i] = vouchers.Voucher{
			VoucherID:  v.VoucherId,
			Discount:   v.Discount,
			ValidUntil: v.ValidUntil,
			Used:       v.Used,
		}
	}
	return vouchersList, nil
}
