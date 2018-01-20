package org.sackfix.field

import org.sackfix.common.validated.fields.SfFieldString

/**
  * Generated by SackFix code generator on 20170404
  */
case class NotifyBrokerOfCreditStringField(override val value: String) extends SfFieldString(208, value) {
  override def toString = appendStringBuilder().toString
  override def appendStringBuilder(b:StringBuilder = new StringBuilder()) = b.append("(208)NotifyBrokerOfCreditString=(").append(value).append(")")
}

object NotifyBrokerOfCreditStringField {
  val TagId = 208  
  def decode(a: Option[Any]) : Option[NotifyBrokerOfCreditStringField] = a match {
    case Some(v) => decode(v)
    case _ =>  scala.Option.empty[NotifyBrokerOfCreditStringField]
  }

  def decode(a: Any) : Option[NotifyBrokerOfCreditStringField] = a match {
    case v: String => Some(NotifyBrokerOfCreditStringField(v))
    case v: NotifyBrokerOfCreditStringField => Some(v)
    case _ => scala.Option.empty[NotifyBrokerOfCreditStringField]
  } 
}