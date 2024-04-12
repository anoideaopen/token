package service

import (
	"context"
	"math/big"
	"reflect"
	"testing"

	"github.com/anoideaopen/token/model"
	"go.uber.org/mock/gomock"
)

func TestBalance_Deposit(t *testing.T) {
	type args struct {
		ctx  context.Context
		addr model.Address
		acc  model.Account
		curr model.Currency
		val  *big.Int
	}
	tests := []struct {
		name    string
		bs      *Balance
		args    args
		want    model.BalanceUpdate
		wantErr bool
	}{
		{
			name: "success",
			bs: func() *Balance {
				env := newEnvironment(t)
				gomock.InOrder(
					env.repoBalance.EXPECT().Load(
						gomock.Any(),
						user1.address,
						user1.account1.account,
						user1.account1.currency,
					).Return(user1.account1.balance, nil),
					env.repoBalance.EXPECT().Save(
						gomock.Any(),
						user1.address,
						user1.account1.account,
						user1.account1.currency,
						big.NewInt(200),
					).Return(nil),
				)

				return &Balance{env.repoBalance}
			}(),
			args: args{
				ctx:  ctx,
				addr: user1.address,
				acc:  user1.account1.account,
				curr: user1.account1.currency,
				val:  big.NewInt(100),
			},
			want: model.BalanceUpdate{
				Address:    user1.address,
				Account:    user1.account1.account,
				Currency:   "USD",
				OldValue:   user1.account1.balance,
				NewValue:   big.NewInt(200),
				ValueDelta: big.NewInt(100),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.bs.Deposit(tt.args.ctx, tt.args.addr, tt.args.acc, tt.args.curr, tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("Balance.Deposit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Balance.Deposit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBalance_Withdraw(t *testing.T) {
	type args struct {
		ctx  context.Context
		addr model.Address
		acc  model.Account
		curr model.Currency
		val  *big.Int
	}
	tests := []struct {
		name    string
		bs      *Balance
		args    args
		want    model.BalanceUpdate
		wantErr bool
	}{
		{
			name: "success",
			bs: func() *Balance {
				env := newEnvironment(t)
				gomock.InOrder(
					env.repoBalance.EXPECT().Load(
						gomock.Any(),
						user2.address,
						user2.account1.account,
						user2.account1.currency,
					).Return(user2.account1.balance, nil),
					env.repoBalance.EXPECT().Save(
						gomock.Any(),
						user2.address,
						user2.account1.account,
						user2.account1.currency,
						big.NewInt(200),
					).Return(nil),
				)

				return &Balance{env.repoBalance}
			}(),
			args: args{
				ctx:  ctx,
				addr: user2.address,
				acc:  user2.account1.account,
				curr: user2.account1.currency,
				val:  big.NewInt(100),
			},
			want: model.BalanceUpdate{
				Address:    user2.address,
				Account:    user2.account1.account,
				Currency:   user2.account1.currency,
				OldValue:   user2.account1.balance,
				NewValue:   big.NewInt(200),
				ValueDelta: big.NewInt(100),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.bs.Withdraw(tt.args.ctx, tt.args.addr, tt.args.acc, tt.args.curr, tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("Balance.Withdraw() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Balance.Withdraw() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBalance_Transfer(t *testing.T) {
	type args struct {
		ctx      context.Context
		addrFrom model.Address
		addrTo   model.Address
		acc      model.Account
		curr     model.Currency
		val      *big.Int
	}
	tests := []struct {
		name    string
		bs      *Balance
		args    args
		want    [2]model.BalanceUpdate
		wantErr bool
	}{
		{
			name: "success",
			bs: func() *Balance {
				env := newEnvironment(t)
				gomock.InOrder(
					env.repoBalance.EXPECT().Load(
						gomock.Any(),
						user1.address,
						user1.account1.account,
						user1.account1.currency,
					).Return(user1.account1.balance, nil),
					env.repoBalance.EXPECT().Load(
						gomock.Any(),
						user2.address,
						user2.account1.account,
						user2.account1.currency,
					).Return(user2.account1.balance, nil),
					env.repoBalance.EXPECT().Save(
						gomock.Any(),
						user1.address,
						user1.account1.account,
						user1.account1.currency,
						big.NewInt(50),
					).Return(nil),
					env.repoBalance.EXPECT().Save(
						gomock.Any(),
						user2.address,
						user2.account1.account,
						user2.account1.currency,
						big.NewInt(350),
					).Return(nil),
				)
				return &Balance{env.repoBalance}
			}(),
			args: args{
				ctx:      ctx,
				addrFrom: user1.address,
				addrTo:   user2.address,
				acc:      user1.account1.account,
				curr:     user1.account1.currency,
				val:      big.NewInt(50),
			},
			want: [2]model.BalanceUpdate{
				{
					Address:    user1.address,
					Account:    user1.account1.account,
					Currency:   user1.account1.currency,
					OldValue:   user1.account1.balance,
					NewValue:   big.NewInt(50),
					ValueDelta: big.NewInt(50),
				},
				{
					Address:    user2.address,
					Account:    user2.account1.account,
					Currency:   user1.account1.currency,
					OldValue:   user2.account1.balance,
					NewValue:   big.NewInt(350),
					ValueDelta: big.NewInt(50),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.bs.Transfer(tt.args.ctx, tt.args.addrFrom, tt.args.addrTo, tt.args.acc, tt.args.curr, tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("Balance.Transfer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Balance.Transfer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBalance_InternalTransfer(t *testing.T) {
	type args struct {
		ctx     context.Context
		addr    model.Address
		accFrom model.Account
		accTo   model.Account
		curr    model.Currency
		val     *big.Int
	}
	tests := []struct {
		name    string
		bs      *Balance
		args    args
		want    [2]model.BalanceUpdate
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.bs.InternalTransfer(tt.args.ctx, tt.args.addr, tt.args.accFrom, tt.args.accTo, tt.args.curr, tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("Balance.InternalTransfer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Balance.InternalTransfer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBalance_transfer(t *testing.T) {
	type args struct {
		ctx      context.Context
		addrFrom model.Address
		addrTo   model.Address
		accFrom  model.Account
		accTo    model.Account
		curr     model.Currency
		amt      *big.Int
	}
	tests := []struct {
		name    string
		bs      *Balance
		args    args
		want    [2]model.BalanceUpdate
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.bs.transfer(tt.args.ctx, tt.args.addrFrom, tt.args.addrTo, tt.args.accFrom, tt.args.accTo, tt.args.curr, tt.args.amt)
			if (err != nil) != tt.wantErr {
				t.Errorf("Balance.transfer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Balance.transfer() = %v, want %v", got, tt.want)
			}
		})
	}
}
