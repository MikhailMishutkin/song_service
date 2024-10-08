/*
 * Time Tracker
 * API for control user's working time
 *
 * The version of the OpenAPI document: 1.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package org.openapitools.client.model;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;

/**
 * ModelsPeriod
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2024-07-05T13:31:14.907388400+03:00[Europe/Moscow]")
public class ModelsPeriod {
  public static final String SERIALIZED_NAME_FIRST_DATE = "firstDate";
  @SerializedName(SERIALIZED_NAME_FIRST_DATE)
  private String firstDate;

  public static final String SERIALIZED_NAME_SECOND_DATE = "secondDate";
  @SerializedName(SERIALIZED_NAME_SECOND_DATE)
  private String secondDate;


  public ModelsPeriod firstDate(String firstDate) {
    
    this.firstDate = firstDate;
    return this;
  }

   /**
   * Get firstDate
   * @return firstDate
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getFirstDate() {
    return firstDate;
  }


  public void setFirstDate(String firstDate) {
    this.firstDate = firstDate;
  }


  public ModelsPeriod secondDate(String secondDate) {
    
    this.secondDate = secondDate;
    return this;
  }

   /**
   * Get secondDate
   * @return secondDate
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getSecondDate() {
    return secondDate;
  }


  public void setSecondDate(String secondDate) {
    this.secondDate = secondDate;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ModelsPeriod modelsPeriod = (ModelsPeriod) o;
    return Objects.equals(this.firstDate, modelsPeriod.firstDate) &&
        Objects.equals(this.secondDate, modelsPeriod.secondDate);
  }

  @Override
  public int hashCode() {
    return Objects.hash(firstDate, secondDate);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ModelsPeriod {\n");
    sb.append("    firstDate: ").append(toIndentedString(firstDate)).append("\n");
    sb.append("    secondDate: ").append(toIndentedString(secondDate)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

}

