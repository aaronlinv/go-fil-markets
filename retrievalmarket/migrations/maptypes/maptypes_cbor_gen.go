// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package maptypes

import (
	"fmt"
	"io"
	"sort"

	piecestore "github.com/filecoin-project/go-fil-markets/piecestore"
	retrievalmarket "github.com/filecoin-project/go-fil-markets/retrievalmarket"
	multistore "github.com/filecoin-project/go-multistore"
	cid "github.com/ipfs/go-cid"
	peer "github.com/libp2p/go-libp2p-core/peer"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

var _ = xerrors.Errorf
var _ = cid.Undef
var _ = sort.Sort

func (t *ClientDealState1) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{181}); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.DealProposal (retrievalmarket.DealProposal) (struct)
	if len("DealProposal") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"DealProposal\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("DealProposal"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("DealProposal")); err != nil {
		return err
	}

	if err := t.DealProposal.MarshalCBOR(w); err != nil {
		return err
	}

	// t.StoreID (multistore.StoreID) (uint64)
	if len("StoreID") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"StoreID\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("StoreID"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("StoreID")); err != nil {
		return err
	}

	if t.StoreID == nil {
		if _, err := w.Write(cbg.CborNull); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(*t.StoreID)); err != nil {
			return err
		}
	}

	// t.ChannelID (datatransfer.ChannelID) (struct)
	if len("ChannelID") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"ChannelID\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("ChannelID"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("ChannelID")); err != nil {
		return err
	}

	if err := t.ChannelID.MarshalCBOR(w); err != nil {
		return err
	}

	// t.LastPaymentRequested (bool) (bool)
	if len("LastPaymentRequested") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"LastPaymentRequested\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("LastPaymentRequested"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("LastPaymentRequested")); err != nil {
		return err
	}

	if err := cbg.WriteBool(w, t.LastPaymentRequested); err != nil {
		return err
	}

	// t.AllBlocksReceived (bool) (bool)
	if len("AllBlocksReceived") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"AllBlocksReceived\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("AllBlocksReceived"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("AllBlocksReceived")); err != nil {
		return err
	}

	if err := cbg.WriteBool(w, t.AllBlocksReceived); err != nil {
		return err
	}

	// t.TotalFunds (big.Int) (struct)
	if len("TotalFunds") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"TotalFunds\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("TotalFunds"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("TotalFunds")); err != nil {
		return err
	}

	if err := t.TotalFunds.MarshalCBOR(w); err != nil {
		return err
	}

	// t.ClientWallet (address.Address) (struct)
	if len("ClientWallet") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"ClientWallet\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("ClientWallet"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("ClientWallet")); err != nil {
		return err
	}

	if err := t.ClientWallet.MarshalCBOR(w); err != nil {
		return err
	}

	// t.MinerWallet (address.Address) (struct)
	if len("MinerWallet") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"MinerWallet\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("MinerWallet"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("MinerWallet")); err != nil {
		return err
	}

	if err := t.MinerWallet.MarshalCBOR(w); err != nil {
		return err
	}

	// t.PaymentInfo (retrievalmarket.PaymentInfo) (struct)
	if len("PaymentInfo") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"PaymentInfo\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("PaymentInfo"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("PaymentInfo")); err != nil {
		return err
	}

	if err := t.PaymentInfo.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Status (retrievalmarket.DealStatus) (uint64)
	if len("Status") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Status\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Status"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Status")); err != nil {
		return err
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.Status)); err != nil {
		return err
	}

	// t.Sender (peer.ID) (string)
	if len("Sender") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Sender\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Sender"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Sender")); err != nil {
		return err
	}

	if len(t.Sender) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Sender was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len(t.Sender))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.Sender)); err != nil {
		return err
	}

	// t.TotalReceived (uint64) (uint64)
	if len("TotalReceived") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"TotalReceived\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("TotalReceived"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("TotalReceived")); err != nil {
		return err
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.TotalReceived)); err != nil {
		return err
	}

	// t.Message (string) (string)
	if len("Message") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Message\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Message"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Message")); err != nil {
		return err
	}

	if len(t.Message) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Message was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len(t.Message))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.Message)); err != nil {
		return err
	}

	// t.BytesPaidFor (uint64) (uint64)
	if len("BytesPaidFor") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"BytesPaidFor\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("BytesPaidFor"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("BytesPaidFor")); err != nil {
		return err
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.BytesPaidFor)); err != nil {
		return err
	}

	// t.CurrentInterval (uint64) (uint64)
	if len("CurrentInterval") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"CurrentInterval\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("CurrentInterval"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("CurrentInterval")); err != nil {
		return err
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.CurrentInterval)); err != nil {
		return err
	}

	// t.PaymentRequested (big.Int) (struct)
	if len("PaymentRequested") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"PaymentRequested\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("PaymentRequested"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("PaymentRequested")); err != nil {
		return err
	}

	if err := t.PaymentRequested.MarshalCBOR(w); err != nil {
		return err
	}

	// t.FundsSpent (big.Int) (struct)
	if len("FundsSpent") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"FundsSpent\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("FundsSpent"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("FundsSpent")); err != nil {
		return err
	}

	if err := t.FundsSpent.MarshalCBOR(w); err != nil {
		return err
	}

	// t.UnsealFundsPaid (big.Int) (struct)
	if len("UnsealFundsPaid") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"UnsealFundsPaid\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("UnsealFundsPaid"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("UnsealFundsPaid")); err != nil {
		return err
	}

	if err := t.UnsealFundsPaid.MarshalCBOR(w); err != nil {
		return err
	}

	// t.WaitMsgCID (cid.Cid) (struct)
	if len("WaitMsgCID") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"WaitMsgCID\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("WaitMsgCID"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("WaitMsgCID")); err != nil {
		return err
	}

	if t.WaitMsgCID == nil {
		if _, err := w.Write(cbg.CborNull); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteCidBuf(scratch, w, *t.WaitMsgCID); err != nil {
			return xerrors.Errorf("failed to write cid field t.WaitMsgCID: %w", err)
		}
	}

	// t.VoucherShortfall (big.Int) (struct)
	if len("VoucherShortfall") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"VoucherShortfall\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("VoucherShortfall"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("VoucherShortfall")); err != nil {
		return err
	}

	if err := t.VoucherShortfall.MarshalCBOR(w); err != nil {
		return err
	}

	// t.LegacyProtocol (bool) (bool)
	if len("LegacyProtocol") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"LegacyProtocol\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("LegacyProtocol"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("LegacyProtocol")); err != nil {
		return err
	}

	if err := cbg.WriteBool(w, t.LegacyProtocol); err != nil {
		return err
	}
	return nil
}

func (t *ClientDealState1) UnmarshalCBOR(r io.Reader) error {
	*t = ClientDealState1{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("ClientDealState1: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadStringBuf(br, scratch)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.DealProposal (retrievalmarket.DealProposal) (struct)
		case "DealProposal":

			{

				if err := t.DealProposal.UnmarshalCBOR(br); err != nil {
					return xerrors.Errorf("unmarshaling t.DealProposal: %w", err)
				}

			}
			// t.StoreID (multistore.StoreID) (uint64)
		case "StoreID":

			{

				b, err := br.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := br.UnreadByte(); err != nil {
						return err
					}
					maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
					if err != nil {
						return err
					}
					if maj != cbg.MajUnsignedInt {
						return fmt.Errorf("wrong type for uint64 field")
					}
					typed := multistore.StoreID(extra)
					t.StoreID = &typed
				}

			}
			// t.ChannelID (datatransfer.ChannelID) (struct)
		case "ChannelID":

			{

				if err := t.ChannelID.UnmarshalCBOR(br); err != nil {
					return xerrors.Errorf("unmarshaling t.ChannelID: %w", err)
				}

			}
			// t.LastPaymentRequested (bool) (bool)
		case "LastPaymentRequested":

			maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
			if err != nil {
				return err
			}
			if maj != cbg.MajOther {
				return fmt.Errorf("booleans must be major type 7")
			}
			switch extra {
			case 20:
				t.LastPaymentRequested = false
			case 21:
				t.LastPaymentRequested = true
			default:
				return fmt.Errorf("booleans are either major type 7, value 20 or 21 (got %d)", extra)
			}
			// t.AllBlocksReceived (bool) (bool)
		case "AllBlocksReceived":

			maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
			if err != nil {
				return err
			}
			if maj != cbg.MajOther {
				return fmt.Errorf("booleans must be major type 7")
			}
			switch extra {
			case 20:
				t.AllBlocksReceived = false
			case 21:
				t.AllBlocksReceived = true
			default:
				return fmt.Errorf("booleans are either major type 7, value 20 or 21 (got %d)", extra)
			}
			// t.TotalFunds (big.Int) (struct)
		case "TotalFunds":

			{

				if err := t.TotalFunds.UnmarshalCBOR(br); err != nil {
					return xerrors.Errorf("unmarshaling t.TotalFunds: %w", err)
				}

			}
			// t.ClientWallet (address.Address) (struct)
		case "ClientWallet":

			{

				if err := t.ClientWallet.UnmarshalCBOR(br); err != nil {
					return xerrors.Errorf("unmarshaling t.ClientWallet: %w", err)
				}

			}
			// t.MinerWallet (address.Address) (struct)
		case "MinerWallet":

			{

				if err := t.MinerWallet.UnmarshalCBOR(br); err != nil {
					return xerrors.Errorf("unmarshaling t.MinerWallet: %w", err)
				}

			}
			// t.PaymentInfo (retrievalmarket.PaymentInfo) (struct)
		case "PaymentInfo":

			{

				b, err := br.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := br.UnreadByte(); err != nil {
						return err
					}
					t.PaymentInfo = new(retrievalmarket.PaymentInfo)
					if err := t.PaymentInfo.UnmarshalCBOR(br); err != nil {
						return xerrors.Errorf("unmarshaling t.PaymentInfo pointer: %w", err)
					}
				}

			}
			// t.Status (retrievalmarket.DealStatus) (uint64)
		case "Status":

			{

				maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
				if err != nil {
					return err
				}
				if maj != cbg.MajUnsignedInt {
					return fmt.Errorf("wrong type for uint64 field")
				}
				t.Status = retrievalmarket.DealStatus(extra)

			}
			// t.Sender (peer.ID) (string)
		case "Sender":

			{
				sval, err := cbg.ReadStringBuf(br, scratch)
				if err != nil {
					return err
				}

				t.Sender = peer.ID(sval)
			}
			// t.TotalReceived (uint64) (uint64)
		case "TotalReceived":

			{

				maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
				if err != nil {
					return err
				}
				if maj != cbg.MajUnsignedInt {
					return fmt.Errorf("wrong type for uint64 field")
				}
				t.TotalReceived = uint64(extra)

			}
			// t.Message (string) (string)
		case "Message":

			{
				sval, err := cbg.ReadStringBuf(br, scratch)
				if err != nil {
					return err
				}

				t.Message = string(sval)
			}
			// t.BytesPaidFor (uint64) (uint64)
		case "BytesPaidFor":

			{

				maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
				if err != nil {
					return err
				}
				if maj != cbg.MajUnsignedInt {
					return fmt.Errorf("wrong type for uint64 field")
				}
				t.BytesPaidFor = uint64(extra)

			}
			// t.CurrentInterval (uint64) (uint64)
		case "CurrentInterval":

			{

				maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
				if err != nil {
					return err
				}
				if maj != cbg.MajUnsignedInt {
					return fmt.Errorf("wrong type for uint64 field")
				}
				t.CurrentInterval = uint64(extra)

			}
			// t.PaymentRequested (big.Int) (struct)
		case "PaymentRequested":

			{

				if err := t.PaymentRequested.UnmarshalCBOR(br); err != nil {
					return xerrors.Errorf("unmarshaling t.PaymentRequested: %w", err)
				}

			}
			// t.FundsSpent (big.Int) (struct)
		case "FundsSpent":

			{

				if err := t.FundsSpent.UnmarshalCBOR(br); err != nil {
					return xerrors.Errorf("unmarshaling t.FundsSpent: %w", err)
				}

			}
			// t.UnsealFundsPaid (big.Int) (struct)
		case "UnsealFundsPaid":

			{

				if err := t.UnsealFundsPaid.UnmarshalCBOR(br); err != nil {
					return xerrors.Errorf("unmarshaling t.UnsealFundsPaid: %w", err)
				}

			}
			// t.WaitMsgCID (cid.Cid) (struct)
		case "WaitMsgCID":

			{

				b, err := br.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := br.UnreadByte(); err != nil {
						return err
					}

					c, err := cbg.ReadCid(br)
					if err != nil {
						return xerrors.Errorf("failed to read cid field t.WaitMsgCID: %w", err)
					}

					t.WaitMsgCID = &c
				}

			}
			// t.VoucherShortfall (big.Int) (struct)
		case "VoucherShortfall":

			{

				if err := t.VoucherShortfall.UnmarshalCBOR(br); err != nil {
					return xerrors.Errorf("unmarshaling t.VoucherShortfall: %w", err)
				}

			}
			// t.LegacyProtocol (bool) (bool)
		case "LegacyProtocol":

			maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
			if err != nil {
				return err
			}
			if maj != cbg.MajOther {
				return fmt.Errorf("booleans must be major type 7")
			}
			switch extra {
			case 20:
				t.LegacyProtocol = false
			case 21:
				t.LegacyProtocol = true
			default:
				return fmt.Errorf("booleans are either major type 7, value 20 or 21 (got %d)", extra)
			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}
func (t *ProviderDealState1) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{171}); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.DealProposal (retrievalmarket.DealProposal) (struct)
	if len("DealProposal") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"DealProposal\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("DealProposal"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("DealProposal")); err != nil {
		return err
	}

	if err := t.DealProposal.MarshalCBOR(w); err != nil {
		return err
	}

	// t.StoreID (multistore.StoreID) (uint64)
	if len("StoreID") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"StoreID\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("StoreID"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("StoreID")); err != nil {
		return err
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.StoreID)); err != nil {
		return err
	}

	// t.ChannelID (datatransfer.ChannelID) (struct)
	if len("ChannelID") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"ChannelID\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("ChannelID"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("ChannelID")); err != nil {
		return err
	}

	if err := t.ChannelID.MarshalCBOR(w); err != nil {
		return err
	}

	// t.PieceInfo (piecestore.PieceInfo) (struct)
	if len("PieceInfo") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"PieceInfo\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("PieceInfo"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("PieceInfo")); err != nil {
		return err
	}

	if err := t.PieceInfo.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Status (retrievalmarket.DealStatus) (uint64)
	if len("Status") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Status\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Status"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Status")); err != nil {
		return err
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.Status)); err != nil {
		return err
	}

	// t.Receiver (peer.ID) (string)
	if len("Receiver") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Receiver\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Receiver"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Receiver")); err != nil {
		return err
	}

	if len(t.Receiver) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Receiver was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len(t.Receiver))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.Receiver)); err != nil {
		return err
	}

	// t.TotalSent (uint64) (uint64)
	if len("TotalSent") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"TotalSent\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("TotalSent"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("TotalSent")); err != nil {
		return err
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.TotalSent)); err != nil {
		return err
	}

	// t.FundsReceived (big.Int) (struct)
	if len("FundsReceived") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"FundsReceived\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("FundsReceived"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("FundsReceived")); err != nil {
		return err
	}

	if err := t.FundsReceived.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Message (string) (string)
	if len("Message") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Message\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Message"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Message")); err != nil {
		return err
	}

	if len(t.Message) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Message was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len(t.Message))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.Message)); err != nil {
		return err
	}

	// t.CurrentInterval (uint64) (uint64)
	if len("CurrentInterval") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"CurrentInterval\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("CurrentInterval"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("CurrentInterval")); err != nil {
		return err
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.CurrentInterval)); err != nil {
		return err
	}

	// t.LegacyProtocol (bool) (bool)
	if len("LegacyProtocol") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"LegacyProtocol\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("LegacyProtocol"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("LegacyProtocol")); err != nil {
		return err
	}

	if err := cbg.WriteBool(w, t.LegacyProtocol); err != nil {
		return err
	}
	return nil
}

func (t *ProviderDealState1) UnmarshalCBOR(r io.Reader) error {
	*t = ProviderDealState1{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("ProviderDealState1: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadStringBuf(br, scratch)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.DealProposal (retrievalmarket.DealProposal) (struct)
		case "DealProposal":

			{

				if err := t.DealProposal.UnmarshalCBOR(br); err != nil {
					return xerrors.Errorf("unmarshaling t.DealProposal: %w", err)
				}

			}
			// t.StoreID (multistore.StoreID) (uint64)
		case "StoreID":

			{

				maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
				if err != nil {
					return err
				}
				if maj != cbg.MajUnsignedInt {
					return fmt.Errorf("wrong type for uint64 field")
				}
				t.StoreID = multistore.StoreID(extra)

			}
			// t.ChannelID (datatransfer.ChannelID) (struct)
		case "ChannelID":

			{

				if err := t.ChannelID.UnmarshalCBOR(br); err != nil {
					return xerrors.Errorf("unmarshaling t.ChannelID: %w", err)
				}

			}
			// t.PieceInfo (piecestore.PieceInfo) (struct)
		case "PieceInfo":

			{

				b, err := br.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := br.UnreadByte(); err != nil {
						return err
					}
					t.PieceInfo = new(piecestore.PieceInfo)
					if err := t.PieceInfo.UnmarshalCBOR(br); err != nil {
						return xerrors.Errorf("unmarshaling t.PieceInfo pointer: %w", err)
					}
				}

			}
			// t.Status (retrievalmarket.DealStatus) (uint64)
		case "Status":

			{

				maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
				if err != nil {
					return err
				}
				if maj != cbg.MajUnsignedInt {
					return fmt.Errorf("wrong type for uint64 field")
				}
				t.Status = retrievalmarket.DealStatus(extra)

			}
			// t.Receiver (peer.ID) (string)
		case "Receiver":

			{
				sval, err := cbg.ReadStringBuf(br, scratch)
				if err != nil {
					return err
				}

				t.Receiver = peer.ID(sval)
			}
			// t.TotalSent (uint64) (uint64)
		case "TotalSent":

			{

				maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
				if err != nil {
					return err
				}
				if maj != cbg.MajUnsignedInt {
					return fmt.Errorf("wrong type for uint64 field")
				}
				t.TotalSent = uint64(extra)

			}
			// t.FundsReceived (big.Int) (struct)
		case "FundsReceived":

			{

				if err := t.FundsReceived.UnmarshalCBOR(br); err != nil {
					return xerrors.Errorf("unmarshaling t.FundsReceived: %w", err)
				}

			}
			// t.Message (string) (string)
		case "Message":

			{
				sval, err := cbg.ReadStringBuf(br, scratch)
				if err != nil {
					return err
				}

				t.Message = string(sval)
			}
			// t.CurrentInterval (uint64) (uint64)
		case "CurrentInterval":

			{

				maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
				if err != nil {
					return err
				}
				if maj != cbg.MajUnsignedInt {
					return fmt.Errorf("wrong type for uint64 field")
				}
				t.CurrentInterval = uint64(extra)

			}
			// t.LegacyProtocol (bool) (bool)
		case "LegacyProtocol":

			maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
			if err != nil {
				return err
			}
			if maj != cbg.MajOther {
				return fmt.Errorf("booleans must be major type 7")
			}
			switch extra {
			case 20:
				t.LegacyProtocol = false
			case 21:
				t.LegacyProtocol = true
			default:
				return fmt.Errorf("booleans are either major type 7, value 20 or 21 (got %d)", extra)
			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}